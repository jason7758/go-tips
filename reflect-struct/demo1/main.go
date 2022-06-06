package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	Name   string
	Role   string
	Salary float64
	Age    *int
}

var i int = 5
var xiaowang = &Employee{
	Name:   "xiaowang",
	Role:   "glory engineer",
	Salary: 0.5,
	Age:    &i,
}

func traverse(target interface{}) {
	sVal := reflect.ValueOf(target)
	sType := reflect.TypeOf(target)
	if sType.Kind() == reflect.Ptr {
		//用Elem()获得实际的value
		sVal = sVal.Elem()
		sType = sType.Elem()
	}
	num := sVal.NumField()
	for i := 0; i < num; i++ {
		f := sType.Field(i)
		val := sVal.Field(i).Interface()

		fmt.Printf("%5s %v = %v\n", f.Name, f.Type, val)
	}
}

func main() {
	traverse(xiaowang)
}
