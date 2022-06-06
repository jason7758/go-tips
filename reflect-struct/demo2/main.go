package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int = 5
	var a *int

	a = &i
	typeOfA := reflect.TypeOf(a)
	valueOfA := reflect.ValueOf(a)
	if typeOfA.Kind() == reflect.Ptr {
		//用Elem()获得实际的value
		valueOfA = valueOfA.Elem()
		typeOfA = typeOfA.Elem()
	}

	fmt.Println(typeOfA.Name(), typeOfA.Kind(), valueOfA.Interface().(int))

	sVal := valueOfA.Interface().(int)
	fmt.Println(sVal)
}
