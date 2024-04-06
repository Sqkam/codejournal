package main

import (
	"github.com/sqkam/codejournal/gen"
)

func main() {
	gen.CreateStruct()
	gen.SwitchCase()
	gen.JudgeNilSlice(nil)
}
