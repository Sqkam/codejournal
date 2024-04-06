package main

import "github.com/sqkam/codejournal/clashmeta"

type People struct {
	Name string
	Age  int64
}

func (p *People) Ch(pCopy People) {
	*p = pCopy
}

func main() {
	clashmeta.ParseTime()
	clashmeta.ChangeReceiverInFunc()
	clashmeta.CheckFileExist("file")
	clashmeta.CreateSlice()
	clashmeta.CreateList()
	clashmeta.GetFormPool()

}
