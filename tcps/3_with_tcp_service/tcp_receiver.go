package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var (
	wg       = sync.WaitGroup{}
	closech  = make(chan int)
	msgGroup = []Message{}
	addr     = "localhost:6379"
)

type Message struct {
	Id   string
	Text string
}

// 创建一个连接
func retNewConn() net.Conn {

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	return conn
}

// 设置后 服务返回一个ok字符
func SetName(cn net.Conn) {

	cn.Write([]byte("SET name jack \n"))
	scanner := bufio.NewScanner(cn)
	for scanner.Scan() {
		txt := scanner.Text()
		msgGroup = append(msgGroup, Message{Id: fmt.Sprintf("M:%v", time.Now().UnixMilli()), Text: txt})
		break
	}
	return
}

// 查询时返回两个，第一个是长度，第二个是内容
func GetName(cn net.Conn) {

	cn.Write([]byte("GET name \n"))
	scanner := bufio.NewScanner(cn)
	i := 0
	for scanner.Scan() {
		i++
		txt := scanner.Text()
		msgGroup = append(msgGroup, Message{Id: fmt.Sprintf("M:%v", time.Now().UnixMilli()), Text: txt})

		if i >= 2 {
			wg.Done()
			closech <- i
			return
		}
	}

}

// 设置和获取, 查看结果并退出
func ConnTcp() {
	cn := retNewConn()
	SetName(cn)
	wg.Add(1)
	go GetName(cn)
	wg.Wait()

	for i, v := range msgGroup {
		fmt.Printf("i:%v, v:%v\n", i, v)

	}

	select {
	case <-closech:
		cn.Close()
		os.Exit(1)
	default:
		fmt.Printf("never do this.")
	}

}

func main() {
	ConnTcp()
}
