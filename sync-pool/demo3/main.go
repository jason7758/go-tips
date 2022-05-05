package main

import (
	"fmt"
	"sync"
)

var intPool = sync.Pool{
	New: func() interface{} {
		b := make([]int, 8)
		return &b
	},
}

func main() {
	//不使用对象池
	for i := 0; i < 3; i++ {
		obj := make([]int, 8)
		obj[i] = i
		fmt.Printf("obj%d: %T %+v\n", i, obj, obj)
	}
	fmt.Println("-------------")
	//使用对象池
	for i := 0; i < 3; i++ {
		obj := intPool.Get().(*[]int)
		(*obj)[i] = i
		fmt.Printf("obj%d: %T %+v\n", i, obj, obj)
		intPool.Put(obj)
	}
}
