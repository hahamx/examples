package main

import "fmt"

type HashTable interface {
	Write(p []byte) (n int, err error)
	NewSum(b []byte) []byte
}

//使用基础类
func NewHashTable() HashTable {
	return &Base{}
}

//使用子类
func NewHashWriter() HashTable {
	return &ChWriter{}
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

func NewWriter() Writer {
	return &ChWriter{}
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func NewReader() Reader {
	return &ChReader{}
}

type error interface {
	Error() string
}

type Base struct{}

func (b *Base) Write(p []byte) (n int, err error) { return 1, nil }

func (b *Base) NewSum(p []byte) []byte { return []byte{1} }

type ChWriter struct{ Base }

func (b *ChWriter) Write(p []byte) (n int, err error) { return 2, nil }

type ChReader struct{ Base }

func (b *ChReader) Read(p []byte) (n int, err error) { return 2, nil }

//读者反应有错误
func (b *ChReader) ReportError() string { return "err" }

func TestMain() {
	b1 := &Base{}
	w1 := NewWriter()
	r1 := NewReader()

	fmt.Println("interface of HashTable:")
	fmt.Println(b1.Write([]byte("hello")))
	fmt.Println(b1.NewSum([]byte{}))

	fmt.Println("interface of Writer:")
	fmt.Println(w1.Write([]byte{}))
	//错误调用
	// fmt.Println(w1.NewSum([]byte{}))

	fmt.Println("interface of Reader:")
	//使用接口只有子类函数,
	fmt.Println(r1.Read([]byte{}))
	//读者不能反应错误
	//fmt.Println(r1.ReportError())

	fmt.Println("struct of Reader:")
	//使用本类可以使用全部函数 和 组合的类的函数
	r2 := ChReader{}
	//读者反应有错误
	fmt.Println(r2.ReportError())
	fmt.Println(r2.NewSum([]byte{}))

	/*
		    interface of HashTable:
			1 <nil>
			[1]
			interface of Writer:
			2 <nil>
			interface of Reader:
			2 <nil>
			struct of Reader:
			err
			[1]
	*/
}
