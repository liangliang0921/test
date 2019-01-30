package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type comparison struct {
	x, y unsafe.Pointer
	reflect.Type
}

// 实现一个自己equal函数，让为nil的slice和map与非nil但是为空的slice和map相等
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

func main() {
	got := strings.Split("a:b:c", ":")
	want := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(got, want) {
		fmt.Println("Equal err")
	}
	// 下面这个if判断就是一个错误的，切片只能和nil进行比较。但是DeepEqual可以用在任意类型上
	//if got == want {
	//	fmt.Println("Equal SUCCESS")
	//}

	// DeepEqual 的不足之处
	var a, b []string = nil, []string{}
	fmt.Println(reflect.DeepEqual(a, b)) // 输出false，此时说明nil和slice不能和非nil值但是空的slice作比较

	var c, d map[string]int = nil, make(map[string]int) // 输出false，和slice一样的原理
	fmt.Println(reflect.DeepEqual(c, d))

	fmt.Println(Equal(a, b))
	fmt.Println(Equal(c, d))
	//fmt.Println(len(a))
	//fmt.Println(len(b))
	//fmt.Println(len(c))
	//fmt.Println(len(d))

	e := int16(12)
	f := int32(12)
	if !Equal(e, f) {
		fmt.Printf("%d == %d", e, f)
	}

	// 结构体的相等判断
	type Test struct {
		a string
		b int64
	}
	test1 := Test{"123", 123}
	test2 := Test{"123", 123}
	fmt.Println(Equal(test1, test2))

}

// 实现一个自己equal函数，让为nil的slice和map与非nil但是为空的slice和map相等
func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// 循环检查，为了确保算法对于有环的数据结构也能正常退出
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true
		}
		seen[c] == true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()
	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true
	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		if x.Pointer() == y.Pointer() {
			return true
		}
		for _, k := range x.MapKeys() {
			xtr := x.MapIndex(k)
			ytr := y.MapIndex(k)
			if xtr.IsValid() || ytr.IsValid() || equal(xtr, ytr, seen) {
				return false
			}
			return true
		}
	case reflect.Struct:
		if x.Pointer() == y.Pointer() {
			return true
		}
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}
