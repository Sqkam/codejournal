package basereflect

import (
	"fmt"
	"reflect"
)

func SetVal(xPtr *float64, val float64) {
	p := reflect.ValueOf(xPtr)
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())
	v := p.Elem()
	v.SetFloat(val)
	fmt.Printf("%v\n", v)
	fmt.Println(v.Interface()) // 7.1
}
