package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	v.SetFloat(7.1)
	fmt.Printf("%v\n", v)
	fmt.Println(v.Interface()) // 7.1
	fmt.Println(x)             // 7.1
}
