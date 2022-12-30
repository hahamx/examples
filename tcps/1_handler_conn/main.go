package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

const (
	StrMessage   = 1       //消息类型 字符
	CloseMessage = "QUIT"  //退出标记
	MaxSize      = 2       //最多存几个信息
	ports        = ":3900" //哪个端口提供服务
)

var (
	Sers         = NewValues() //默认服务
	Lock         sync.RWMutex
	ErrCloseSent = fmt.Errorf("service: close sent")
)

/*
主服务结构体
*/
type ValueSer struct {
	Service net.Listener
	Fields  map[string]string
	Size    int
	Running bool
}

func NewValues() *ValueSer {
	return &ValueSer{Fields: make(map[string]string, MaxSize), Size: MaxSize}
}

/*
处理主服务存单存入
@param mk mv 需要存入的单子
*/
func (va *ValueSer) PutIn(mk, mv string) bool {

	Lock.Lock()
	defer Lock.Unlock()
	if len(va.Fields) >= va.Size {
		return false
	}

	va.Fields[mk] = mv
	return true
}

/*
处理主服务存单存入
@param mk  需要取出的单子
*/
func (va *ValueSer) PopOut(mk string) string {
	Lock.Lock()
	defer Lock.Unlock()

	if len(va.Fields) <= 0 {
		return ""
	}
	val := va.Fields[mk]
	delete(va.Fields, mk)
	return val
}

/*
处理主服务启动
*/
func (va *ValueSer) Start() net.Listener {

	ser, err := net.Listen("tcp", ports)

	if err != nil {
		log.Fatalln(err)
	}
	va.Running = true
	return ser
}

/*
主体服务处理函数
*/
func (va *ValueSer) TcpServer(val chan Infos) {

	for info := range val {
		if len(info.Cmd) < 1 {
			info.Result <- "need arg info"
			continue
		}

		fmt.Println("Set info", info)
		for kk, ff := range info.Cmd {
			if kk == "POP" {
				value := va.PopOut(ff[0])
				info.Result <- value
			} else if kk == "PUT" {

				k := ff[0]
				v := ff[1]
				value := va.PutIn(k, v)
				info.Result <- fmt.Sprintf("%v", value)
			} else if kk == CloseMessage {

				va.Service.Close()
				os.Exit(1)
			} else {
				info.Result <- "INVALID COMMAND " + kk + "\n"
			}

		}

	}
}

/*
使用者的指令管理结构体
*/
type Infos struct {
	Cmd    map[string][]string
	Result chan string
}

/*
使用者的指令处理函数，负责把结果写回链接中
*/
func handler(infos chan Infos, cn net.Conn) {

	defer cn.Close()

	scer := bufio.NewScanner(cn)
	for scer.Scan() {
		strs := scer.Text()
		if strs == "" {
			msg := fmt.Sprintf(`useage:
						PUT name jack
						POP name`)
			io.WriteString(cn, msg)
			continue
		}
		fs := strings.Fields(strs)

		result := make(chan string)
		infos <- Infos{
			Cmd:    map[string][]string{fs[0]: fs[1:]},
			Result: result,
		}

		io.WriteString(cn, <-result+"\n")
	}

}

/*
服务入口，集成启动函数
*/
func Start() {

	ser := Sers.Start()

	infos := make(chan Infos)
	go Sers.TcpServer(infos)

	for Sers.Running {

		cn, err := ser.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handler(infos, cn)
	}

}

func main() {

	fmt.Println("start service at:", ports)
	Start()
}
