package main

import (
	"fmt"
	"sync"
	"time"
)

var pool *sync.Pool

type Person struct {
	Say  func()
	Name []string
}

func (ps *Person) RegSay(f func()) {
	ps.Say = f
}

var count int

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			count++
			fmt.Println("Creating a new Person")
			return new(Person)
		},
	}
}

func main() {
	initPool()
	pool.Put(&Person{})
	for i := 0; i < 10; i++ {
		go func(index int) {
			//说点什么
			person := pool.Get().(*Person)
			person.RegSay(func() {
				fmt.Println("我是任务: ", index)
			})
			person.Name = append(person.Name, fmt.Sprintf("%d", index))
			person.Say()
			fmt.Println("my Records is ", person.Name)
			pool.Put(person)
		}(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println(count)
}
