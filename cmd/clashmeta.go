package main

import (
	"container/list"
	"context"
	"fmt"
	"os"
	"path"
	"sync"
	"time"
)

type People struct {
	Name string
	Age  int64
}

func (p *People) Ch(pCopy People) {
	*p = pCopy
}

func main() {
	//parseTime()
	//changeReceiverInFunc()
	//checkFileExist("file")
	//createSlice()
	//createList()
	getFormPool()

}

func checkFileExist(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		cwd, _ := os.Getwd()
		defaultUIpath := path.Join(cwd, "ui")
		fmt.Printf("%+v\n", defaultUIpath)
	}

}

func changeReceiverInFunc() {
	p1 := &People{
		Name: "1",
		Age:  2,
	}
	p1.Ch(People{
		Name: "2",
		Age:  3,
	})
	fmt.Printf("%+v\n", p1)

}
func parseTime() {
	var KeepAliveInterval = 15 * time.Second
	KeepAliveInterval = time.Duration(2) * time.Second
	fmt.Printf("%+v\n", KeepAliveInterval)

}
func createSlice() {
	var proxyList []string //nil
	proxyList = append(proxyList, "DIRECT", "REJECT")
	fmt.Printf("%+v\n", proxyList)
}
func createList() {
	ls := list.New()
	ls.PushBack(People{
		Name: "1",
		Age:  2,
	})
	fmt.Printf("%+v\n", ls.Back().Value)
}
func getFormPool() {
	type udpMessage struct {
		Now int64
	}
	var udpMessagePool = sync.Pool{
		New: func() interface{} {
			msg := new(udpMessage)
			msg.Now = time.Now().UnixMilli()
			return msg
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
LOOP:
	for i := 0; i < 100000000000; i++ {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			go func() {
				message := udpMessagePool.Get().(*udpMessage)
				_ = message
				//fmt.Printf("%v\n", message)
			}()
		}

	}
}
