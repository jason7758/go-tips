package main

import (
	"fmt"
	"sync"
	"time"
)

// 一个[]int的对象池, 每个对象为一个[]int
var intPool = sync.Pool{
	New: func() interface{} {
		b := make([]int, 10)
		return &b
	},
}

func main() {
	fmt.Println("====A=====")
	for i := 0; i < 8; i++ {
		go func(i int) {
			obj := intPool.Get().(*[]int)
			(*obj)[i] = i
			fmt.Printf("%d obj: %T %v \n", i, obj, obj)
			intPool.Put(obj)
		}(i)
	}
	time.Sleep(3 * time.Second)
	fmt.Println("=======B=======")
	for i := 0; i < 8; i++ {
		obj := intPool.Get().(*[]int)
		fmt.Printf("%d obj: %T %+v\n", i, obj, obj)
	}

}
