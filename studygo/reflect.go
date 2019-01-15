package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Id    int
	Class int
}

func (s student) Printstruct() {
	fmt.Println("haha")
}

func (s student) ReflectCallFunc(name string, id int) {
	fmt.Println("ReflectCallFunc name is: ", name, " Id is: ", id)
}

func DoFiledAndMethod(input interface{}) {
	// 得到类型的元数据，通过t我们能获取类型定义里面的所有元素
	t := reflect.TypeOf(input)
	// 得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
	v := reflect.ValueOf(input)

	fmt.Println("TypeOf is: ", t)
	fmt.Println("ValueOf is: ", v)
	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("%s:%v=%v\n", field.Name, field.Type, value)
	}

	// 获取方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

// 利用的reflect反射进行方法的调用
func ReflectCallMethod(input interface{}) {
	v := reflect.ValueOf(input)
	methodName := v.MethodByName("ReflectCallFunc")
	args := []reflect.Value{reflect.ValueOf("wudebao"), reflect.ValueOf(30)}
	methodName.Call(args)

	methodName = v.MethodByName("Printstruct")
	args = make([]reflect.Value, 0)
	methodName.Call(args)
}

// reflect 反射包的使用
func main() {
	// 1.首先转化成 reflect 对象
	s := student{"liang", 15, 15}
	DoFiledAndMethod(s)
	ReflectCallMethod(s)
}
