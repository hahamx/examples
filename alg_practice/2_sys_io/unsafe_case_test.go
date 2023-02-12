package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"unsafe"
)

var (
	i      int = 5
	a          = [10]int{}
	sbyte      = [10]byte{}
	ss         = a[:]
	logger     = log.New(os.Stdout, "info-", 18)
	f      FuncFoo
	zhs    = "文"

	preValue = map[string]uintptr{
		"i":       8,
		"a":       80,
		"ss":      24,
		"f":       48,
		"f.c":     10,
		"int_nil": 8,
	}
)

type FuncFoo struct {
	a int
	b string
	c [10]byte
	d float64
}

func setUp(funcName string) func() {
	return func() { fmt.Println("setUp for func:", funcName) }
}

func ErrorHandler(info string, t *testing.T) {
	t.Error(info)
}

func InfoHandler(info string, t *testing.T) {
	logger.Println(info)
	t.Log(info)
}

func FatalHandler(info string, t *testing.T) {
	t.Fatal(info)
}
func TestAlignof(t *testing.T) {

	defer setUp(t.Name())()
	fmt.Printf("\tExecute test:%v\n", t.Name())

	var x int

	b := uintptr(unsafe.Pointer(&x))%unsafe.Alignof(x) == 0
	t.Log("alignof:", b)

	if unsafe.Alignof(i) != preValue["i"] {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), preValue["int_nil"]), t)

	}

	if unsafe.Alignof(a) != preValue["i"] {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), preValue["int_nil"]), t)

	}

	if unsafe.Alignof(ss) != preValue["i"] {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), preValue["int_nil"]), t)

	}

	if unsafe.Alignof(f.a) != preValue["i"] {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), preValue["int_nil"]), t)

	}

	if unsafe.Alignof(f) != preValue["i"] {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), preValue["int_nil"]), t)

	}

	//中文对齐系数 为 8
	if unsafe.Alignof(zhs) != preValue["i"] {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), preValue["i"]), t)
	}

	//空结构体对齐系数 1
	if unsafe.Alignof(struct{}{}) != 1 {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), 1), t)
	}

	//byte 数组对齐系数为 1
	if unsafe.Alignof(sbyte) != 1 {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), 1), t)
	}

	//长度为0 的数组，与其元素的对齐系数相同
	if unsafe.Alignof([0]int{}) != 8 {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), 8), t)
	}

	//长度为0 的数组，与其元素的对齐系数相同
	if unsafe.Alignof([0]struct{}{}) != 1 {
		ErrorHandler(fmt.Sprintf("Alignof: %v not equal %v", unsafe.Sizeof(i), 1), t)
	}

}

type RandomResource struct {
	Cloud               string // 16 bytes
	Name                string // 16 bytes
	HaveDSL             bool   //  1 byte
	PluginVersion       string // 16 bytes
	IsVersionControlled bool   //  1 byte
	TerraformVersion    string // 16 bytes
	ModuleVersionMajor  int32  //  4 bytes
}

type OrderResource struct {
	ModuleVersionMajor  int32  //  4 bytes
	HaveDSL             bool   //  1 byte
	IsVersionControlled bool   //  1 byte
	Cloud               string // 16 bytes
	Name                string // 16 bytes
	PluginVersion       string // 16 bytes
	TerraformVersion    string // 16 bytes

}

// 存储字段 使用的空间与 类型有关，与字段值没有关系
func TestCheckStruct(m *testing.T) {

	var d RandomResource

	d.Cloud = "aws-singapore"
	d.Name = "ec2"
	d.HaveDSL = true
	d.PluginVersion = "3.64"
	d.TerraformVersion = "1.1"
	d.ModuleVersionMajor = 1
	d.IsVersionControlled = true
	m.Log("==============================================================")
	InfoHandler(fmt.Sprintf("随机顺序属性的结构体内存 总共占用 StructType: %T => [%d]\n", d, unsafe.Sizeof(d)), m)
	m.Log("==============================================================")
	InfoHandler(fmt.Sprintf("字段 Cloud :d.Cloud %T => [%d]\n", d.Cloud, unsafe.Sizeof(d.Cloud)), m)
	InfoHandler(fmt.Sprintf("字段 Name :d.Name %T => [%d]\n", d.Name, unsafe.Sizeof(d.Name)), m)
	m.Logf("字段 :d.HaveDSL %T => [%d]\n", d.HaveDSL, unsafe.Sizeof(d.HaveDSL))
	m.Logf("字段 :d.PluginVersion %T => [%d]\n", d.PluginVersion, unsafe.Sizeof(d.PluginVersion))
	m.Logf("字段 :d.IsVersionControlled %T => [%d]\n", d.IsVersionControlled, unsafe.Sizeof(d.IsVersionControlled))
	m.Logf("字段 :d.TerraformVersion %T => [%d]\n", d.TerraformVersion, unsafe.Sizeof(d.TerraformVersion))
	m.Logf("字段 :d.ModuleVersionMajor %T => [%d]\n", d.ModuleVersionMajor, unsafe.Sizeof(d.ModuleVersionMajor))

	var te = OrderResource{}
	te.Cloud = "aws-singapore" 
	te.Name = "ec2"
	te.PluginVersion = "3.64"
	te.TerraformVersion = "1.1"
	te.ModuleVersionMajor = 1
	te.IsVersionControlled = true
	te.HaveDSL = true

	m.Log("==============================================================")
	m.Logf("属性对齐的结构体内存 总共占用  StructType:d %T => [%d]\n", te, unsafe.Sizeof(te))
	m.Log("==============================================================")
	m.Logf("字段 Cloud :d.Cloud %T => [%d]\n", te.Cloud, unsafe.Sizeof(te.Cloud))
	m.Logf("字段 Name :d.Name %T => [%d]\n", te.Name, unsafe.Sizeof(te.Name))
	m.Logf("字段 :d.PluginVersion %T => [%d]\n", te.PluginVersion, unsafe.Sizeof(te.PluginVersion))
	m.Logf("字段 :d.TerraformVersion %T => [%d]\n", te.TerraformVersion, unsafe.Sizeof(te.TerraformVersion))

	m.Logf("字段 :d.ModuleVersionMajor %T => [%d]\n", te.ModuleVersionMajor, unsafe.Sizeof(te.ModuleVersionMajor))
	m.Logf("字段ModuleVersionMajor :d.IsVersionControlled %T => [%d]\n", te.IsVersionControlled, unsafe.Sizeof(te.IsVersionControlled))
	m.Logf("字段 :d.HaveDSL %T => [%d]\n", te.HaveDSL, unsafe.Sizeof(te.HaveDSL))

	te2 := te
	te2.Cloud = "ali2"

	m.Log("改变 te3 将同时改变 te,te3 指向了 te的地址")
	m.Log("复制了对齐结构体，并重新赋值，用于查看字段长度。")
	m.Log("(*te).Cloud:", (te).Cloud, "*te.Cloud", te.Cloud, "te size:", unsafe.Sizeof(te.Cloud), "te Alignof:", unsafe.Alignof(te.Cloud), "te value len:", len(te.Cloud),
		"reflect Align len and field Align len:", reflect.TypeOf(te.Cloud).Align(), reflect.TypeOf(te.Cloud).FieldAlign())
	m.Log("(*te2).Cloud:", (te2).Cloud, "*te2.Cloud", te.Cloud, "te2 size:", unsafe.Sizeof(te2.Cloud), "te2 Alignof:", unsafe.Alignof(te2.Cloud), "te2 value len:", len(te2.Cloud),
		"reflect Align len and field Align len:", reflect.TypeOf(te2.Cloud).Align(), reflect.TypeOf(te.Cloud).FieldAlign())

	te3 := &te
	te3.Cloud = "HWCloud2-asia-southeast-from-big-plant-place-air-local-video-service-picture-merge-from-other-all-company"

	m.Log("(*te3).Cloud:", (*te3).Cloud, "*te3.Cloud", te3.Cloud, "te3 size:", unsafe.Sizeof(te3.Cloud), "te3 Alignof:", unsafe.Alignof(te3.Cloud), "te3 value len:", len(te3.Cloud),
		"reflect Align len and field Align len:", reflect.TypeOf(te3.Cloud).Align(), reflect.TypeOf(te.Cloud).FieldAlign())

	m.Logf("结构体1字段 Cloud:%v te2:%p\n", (te).Cloud, &te)
	m.Logf("结构体2字段 Cloud:%v te2:%p\n", (te2).Cloud, &te2)
	m.Logf("结构体3字段 Cloud:%v te3:%p\n", (*te3).Cloud, te3)
	m.Logf("字段 Cloud:%v order:%v te:%v, addr:%p\n", te.Cloud, (te).Cloud, te, &te)
 }

 