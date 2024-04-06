package main

import (
	"fmt"
	"github.com/sqkam/codejournal/basereflect"
)

func main() {
	var x float64 = 3.4
	basereflect.SetVal(&x, 7.1)
	fmt.Println(x) // 7.1

}
