package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"os"
)

var ps *process.Process

// 输出内存状态信息
func mem(n int) {
	if ps == nil {
		p, err := process.NewProcess(int32(os.Getpid()))
		if err != nil {
			panic(err)
		}
		ps = p
	}
	mem, _ := ps.MemoryInfo()
	fmt.Printf("%d. VMS: %d MB, RSS: %d MB\n", n, mem.VMS>>20, mem.RSS>>20)
}

func main() {
	// 初始化结束后的内存状态
	mem(1)

	// 创建一个 10 * 1MB 数组后的内存状态
	data := new([10][1024 * 1024]byte)
	mem(2)

	// 填充该数组过程中的内存状态
	for i := range data {
		for x, n := 0, len(data[i]); x < n; x++ {
			data[i][x] = 1
		}
		mem(3)
	}
}
