package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"unicode"
)

const (
	CloseMessage = 'Q'
	Ports        = ":8910"
)

var (
	ErrCloseSent = fmt.Errorf("service: close sent")
	TouchChar    = map[string]bool{
		"0":  true,
		"1":  true,
		"2":  true,
		"3":  true,
		"4":  true,
		"5":  true,
		"6":  true,
		"7":  true,
		"8":  true,
		"9":  true,
		":":  true,
		";":  true,
		"<":  true,
		"=":  true,
		">":  true,
		"?":  true,
		"@":  true,
		"A":  true,
		"B":  true,
		"C":  true,
		"D":  true,
		"E":  true,
		"F":  true,
		"G":  true,
		"H":  true,
		"I":  true,
		"J":  true,
		"K":  true,
		"L":  true,
		"M":  true,
		"N":  true,
		"O":  true,
		"P":  true,
		"Q":  true,
		"R":  true,
		"S":  true,
		"T":  true,
		"U":  true,
		"V":  true,
		"W":  true,
		"X":  true,
		"Y":  true,
		"Z":  true,
		"[":  true,
		"\\": true,
		"]":  true,
		"_":  true,
		"`":  true,
		"a":  true,
		"b":  true,
		"c":  true,
		"d":  true,
		"e":  true,
		"f":  true,
		"g":  true,
		"h":  true,
		"i":  true,
		"j":  true,
		"k":  true,
		"l":  true,
		"m":  true,
		"n":  true,
		"o":  true,
		"p":  true,
		"q":  true,
		"r":  true,
		"s":  true,
		"t":  true,
		"u":  true,
		"v":  true,
		"w":  true,
		"x":  true,
		"y":  true,
		"z":  true,
	}
)

/*
@param 任务执行函数
*/
func do_jobs(v byte, cn net.Conn) []byte {
	if !TouchChar[string(v)] {
		return nil
	}
	retMsg := []byte(fmt.Sprintf("\nDO: %v Addr:%v Result:\n", string(v), cn.RemoteAddr()))
	return retMsg
}

/*
@param cn 待处理的连接
*/
func handler(cn net.Listener) {
	defer cn.Close()

	for {
		conn, err := cn.Accept()
		if err != nil {
			panic(err)
		}
		conn.Write([]byte(fmt.Sprintf("enter something:")))

		for {
			bs := make([]byte, 1024)

			n, err := conn.Read(bs)
			if err != nil {
				break
			}
			if n <= 0 {
				continue
			}
			for _, v := range bs[:n] {
				fmt.Println("receive cahr:", string(rune(v)))
				if len(string(bs[:n])) <= 0 {
					continue
				} else if string(v) == " " || string(v) == "," {
					continue
				} else if string(v) == "\t" {
					continue
				} else if string(v) == "\r" || string(v) == "\n" {
					continue
				} else if v == CloseMessage {
					retMsg := []byte(fmt.Sprintf("%v: %v\n", ErrCloseSent, string(v)))
					conn.Write(retMsg)
					os.Exit(1)
				} else if unicode.IsLetter(rune(v)) {
					result := do_jobs(v, conn)
					_, err = conn.Write(result)
					if err != nil {
						break
					}

				} else {
					fmt.Printf("Client Say:%v \n", string(v))
				}

			}

			spaces := strings.ReplaceAll(string(bs[:n]), " ", "")
			retMsg := []byte(fmt.Sprintf("%v%v\n", spaces, spaces))
			_, err = conn.Write(retMsg)
			if err != nil {
				break
			}
			continue

		}

		conn.Close()
	}
}

/*
服务管理函数
*/
func TcpStart() {
	cn, err := net.Listen("tcp", Ports)
	if err != nil {
		panic(err)
	}
	defer cn.Close()

	handler(cn)
}

func main() {
	fmt.Println("start service:", Ports)
	TcpStart()
}
