package main

import "C"
import (
	"container/list"
	"context"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net"
	"os"
	"path"
	"sync"
	"time"
)

type Kk struct {
	Name string
	Age  int64
}

func (kk *Kk) Ch(kkCopy Kk) {
	*kk = kkCopy
}

// sed 's#.*define.*SRC_V6.*#define SRC_V6=kk#g' /etc/nftables.conf
var cc = make(map[string]string)

func main() {
	//parseTime()
	//changeReceiverInFunc()
	//checkFileExist("kk")
	//createSlice()
	//createList()
	//ip, err := GetOutboundIP()
	//fmt.Println(ip, err)
	//getIP()
	//getFormPool()
	//gormUpdate()

}
func gormUpdate() {
	type Product struct {
		gorm.Model
		Code  string
		Price uint
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	db = db.Debug()
	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	//db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}

func checkFileExist(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		cwd, _ := os.Getwd()
		defaultUIpath := path.Join(cwd, "ui")
		fmt.Printf("%+v\n", defaultUIpath)
	}

}

func changeReceiverInFunc() {
	//  方法中 改变接受者
	k1 := &Kk{
		Name: "1",
		Age:  2,
	}
	k1.Ch(Kk{
		Name: "2",
		Age:  3,
	})
	fmt.Printf("%+v\n", k1)

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
	ls.PushBack(Kk{
		Name: "1",
		Age:  2,
	})
	fmt.Printf("%+v\n", ls.Back().Value)
}

func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "[2001:db8::1]:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func getIP() {

	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()

		fmt.Println(i.Name)
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Printf("%+v\n", ip)
		}
	}

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
	for i := 0; i < 10000000000; i++ {
		go func() {
			message := udpMessagePool.Get().(*udpMessage)
			fmt.Printf("%+v\n", message.Now)
		}()

	}
	<-ctx.Done()

}
