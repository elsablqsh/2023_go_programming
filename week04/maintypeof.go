package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 100
	var b = 4.5
	var c = "Hello!!"

	fmt.Println(a, b, c)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf('x'))
	fmt.Println(reflect.TypeOf(4.3))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf("Elsa"))

}
