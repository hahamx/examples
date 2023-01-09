package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var (
	wg       = sync.WaitGroup{}
	closech  = make(chan int)
	msgGroup = []Message{}
	logs     = log.New(os.Stdout, "INFO-", 13)
	addr     = "192.168.30.131:6379"
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
			// closech <- i
			return
		}
	}

}

// 查询订阅时返回两个，第一个是长度，第二个是内容, 接受20次
func SubScriPtion(cn net.Conn) {
	subcmd := "SUBSCRIBE boards:zoo:visits \n"
	cn.Write([]byte(subcmd))
	scanner := bufio.NewScanner(cn)
	i := 0
	for scanner.Scan() {
		txt := scanner.Text()
		ids := fmt.Sprintf("M:%v", time.Now().UnixMilli())
		nsg := Message{Id: ids, Text: txt}

		if len(txt) >= 25 {
			i++
			logs.Printf("msg length:%v msg:%#v\n", len(txt), nsg)
		}

		msgGroup = append(msgGroup)

		if i >= 200 {
			wg.Done()
			closech <- i
			return
		}
	}
}

// 设置和获取, 查看结果并退出
func ConnTcp() {
	cn := retNewConn()
	// SetName(cn)
	fmt.Printf("start redis of:%v, conn local:%v\n", addr, cn.LocalAddr())

	wg.Add(1)
	// go GetName(cn)
	go SubScriPtion(cn)
	wg.Wait()

	for i, v := range msgGroup {
		fmt.Printf("msg of:%v, txt:%v\n", i, v)

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
