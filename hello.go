package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func main() {

	fmt.Println("hello world")
	fmt.Println("randomInt", rand.Intn(10))
	fmt.Printf("a %g b\n", math.Sqrt(9))
	// 大写开头的为导出的
	fmt.Println(math.Pi)
	fmt.Println(add(5, 2))
	fmt.Println(swap("x", "y"))
	fmt.Println(split(10))
	fmt.Println(a, b, c)
	fmt.Println(a1, b1, c1)
	assignment()
	basicTypes()
	typeConvert()
	typeDerive()
	constTest()
	numConstantTest()
	forTest()
	ifTest()
	switchTest()
	deferTest()
	deferTest2()
	pointerTest()
	structTest()
	arrayTest()
	sliceTest()
}

// 函数结构
func add(x, y int) int {
	return x + y
}

// 函数多值返回
func swap(a, b string) (string, string) {
	return b, a
}

// 没有参数的return返回已命名的返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 声明变量列表
var a, b, c bool

// 声明同时赋值
var a1, b1, c1 = true, false, "c1"

// :=在函数体内可代替var声明
func assignment() {
	var a = 1
	b := 2
	c, d, e := true, false, "e"
	fmt.Println(a, b, c, d, e)
}

// 基本类型
// bool
// string
// int int8 int16 int32(rune) int64
// uint uint(byte) uint16 uint32 uint64 uintptr
// float32 float64
// complex64 complex128
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(15 + 12i)
)

func basicTypes() {
	fmt.Printf("Type:%T Value:%v\n", ToBe, ToBe)
	fmt.Printf("Type:%T Value:%v\n", MaxInt, MaxInt)
	fmt.Printf("Type:%T Value:%v\n", z, z)
}

// GO需要显式类型转换
func typeConvert() {
	a := 1
	b := 2.5
	fmt.Printf("Type:%T Value:%v\n", float64(a), float64(a))
	fmt.Printf("Type:%T Value:%v\n", uint(b), uint(b))
}

// 类型推导
func typeDerive() {
	var i int
	j := i
	fmt.Printf("Type:%T\n", j)
	// 当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度
	b := 0.5i + 0.123
	fmt.Printf("Type:%T\n", b)
}

// const不能使用:=
const consta = 1
const constb = false
const constc = "c"

func constTest() {
	fmt.Println(consta, constb, constc)
}

// 数值常量是高精度的值
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}
func numConstantTest() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func forTest() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum", sum)
	sum1 := 1
	for ; sum1 < 100; {
		sum1 += sum1
	}
	fmt.Println("sum1", sum1)
	sum2 := 1
	// 相当于while
	for sum2 < 100 {
		sum2 += sum2
	}
	fmt.Println("sum2", sum2)
	// 无限循环
	//for {
	//	fmt.Println("无限")
	//}
}

func ifTest() {
	b := 2
	i := 2 * b
	if i < 5 {
		fmt.Println("i<5", i < 5)
	}
	// if的简短语句，作用域if体内
	if j := math.Pow(5, 2); j >= 25 {
		fmt.Println("j>=25", true)
	} else {
		fmt.Println("else")
	}
}
func switchTest() {
	// Go 自动提供了在这些语言中每个 case 后面所需的 break 语句
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux.")
	default:
		fmt.Printf("%s. \n", os)
	}

	// switch 的求值顺序 从上到下顺次执行，直到匹配成功时停止。
	today := time.Now().Weekday()
	fmt.Printf("Today is %s, when is Saturday?\n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Two far away.")
	}

	// 没有条件的 switch 同 switch true 一样。
	// 这种形式能将一长串 if-then-else 写得更加清晰。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferTest() {
	// defer 语句会将函数推迟到外层函数返回之后执行。
	defer fmt.Println("world")
	fmt.Println("hello")
}

func deferTest2() {
	// 推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
	fmt.Println("Counting")
	fmt.Println("Done")
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

func pointerTest() {
	// int类型指针
	var p *int
	i := 2
	// & 操作符会生成一个指向其操作数的指针。
	p = &i
	// 打印内存地址
	fmt.Println(p)
	// 打印值（通过指针读取i）
	fmt.Println(*p)
	// 通过指针设置i
	*p = 3
	fmt.Println(i)
}

// 结构体（struct）就是一组字段
type Coordinate struct {
	X int
	Y int
}

func structTest() {
	v := Coordinate{1, 2}
	// 结构体字段使用点号访问
	fmt.Println(v.X, v.Y)
	// 如果我们有一个指向结构体的指针 p，那么可以通过 (*p).X 来访问其字段 X。
	//不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	p := &v
	(*p).X = 3
	fmt.Println((*p).X)
	p.X = 4
	fmt.Println(p.X)

	var (
		v1 = Coordinate{1, 2}
		// Y隐式0
		v2 = Coordinate{X: 1}
		// X,Y隐式0
		v3 = Coordinate{}
		p1 = &Coordinate{1, 2}
	)
	fmt.Println(v1, v2, v3, p1)
}

func arrayTest() {
	// 数组数量固定
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	numArray := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(numArray)
}

func sliceTest() {
	// 切片
	// 数组大小固定，切片动态可变，更常用
	// 切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：
	// a[low,high]
	// 包含low,不包含high
	arrA := [3]int{1, 2, 3}
	sliceA := arrA[0:2]
	fmt.Println(sliceA)
	// 切片并不存储任何数据，它只是描述了底层数组中的一段。
	// 更改切片的元素会修改其底层数组中对应的元素。
	// 与它共享底层数组的切片都会观测到这些修改。
	names := [4]string{"john", "lihua", "meimei", "xiaoming"}
	fmt.Println(names)
	namesa := names[0:2]
	namesb := names[1:4]
	fmt.Println(namesa, namesb)
	names[1] = "aaa"
	fmt.Println(names, namesa, namesb)
	// 数组
	arraya1 := [2]int{1, 2}
	// 创建数组并构建了一个引用它的切片
	slicea1 := []int{1, 2}
	fmt.Println(arraya1, slicea1)
	// 结构体切片
	sliceStruct := []struct {
		name string
		age  int
	}{
		{"xiaohong", 18},
		{"lihua", 19},
		{"hanmeimei", 20},
	}
	fmt.Println(sliceStruct)
	// 切片默认行为
	// 在进行切片时，你可以利用它的默认行为来忽略上下界。
	// 切片下界的默认值为 0，上界则是该切片的长度。
	arrayb1 := [5]int{1, 2, 3, 4, 5}
	sliceb1 := arrayb1[0:5]
	sliceb2 := arrayb1[:5]
	sliceb3 := arrayb1[0:]
	sliceb4 := arrayb1[:]
	fmt.Println(sliceb1, sliceb2, sliceb3, sliceb4)
	// 切片的长度：包含的元素个数
	// 切片的容量：底层数组的元素个数
	arrayc1 := [5]int{1, 2, 3, 4, 5}
	slicec1 := arrayc1[:0]
	fmt.Printf("len: %d cap: %d\n", len(slicec1), cap(slicec1))
	slicec1 = arrayc1[2:]
	fmt.Printf("len: %d cap: %d\n", len(slicec1), cap(slicec1))
	// 切片的零值是nil
	var sliced1 []int
	fmt.Printf("nil slice-len: %d cap: %d\n", len(sliced1), cap(sliced1))
	fmt.Println(sliced1)
	if sliced1 == nil {
		fmt.Println("nil!")
	}
	// 使用make创建切片
	slicef1 := make([]int, 5)
	fmt.Printf("slicef1 %v len=%d cap=%d\n", slicef1, len(slicef1), cap(slicef1))
	slicef2 := make([]int, 0, 5)
	fmt.Printf("slicef2 %v len=%d cap=%d\n", slicef2, len(slicef2), cap(slicef2))
	slicef3 := slicef2[:2]
	fmt.Printf("slicef3 %v len=%d cap=%d\n", slicef3, len(slicef3), cap(slicef3))
	slicef4 := slicef2[2:5]
	fmt.Printf("slicef4 %v len=%d cap=%d\n", slicef4, len(slicef4), cap(slicef4))
	// 切片的切片
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	// 向切片追加元素
	var sliceh1 []int
	printSlice("sliceh1", sliceh1)
	sliceh1 = append(sliceh1, 0)
	printSlice("sliceh1", sliceh1)
	sliceh1 = append(sliceh1, 1, 2, 3)
	printSlice("sliceh1", sliceh1)
	// for+range循环
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d=%v\n", i, v)
	}
	// range _忽略
	for i, _ := range pow {
		fmt.Println(i)
	}
	fmt.Println("----------------")
	for _, v := range pow {
		fmt.Println(v)
	}
	// map映射
	var map1 map[string]Point
	map1 = make(map[string]Point)
	map1["pointa"] = Point{1, 2}
	map1["pointb"] = Point{3, 2}
	fmt.Println(map1["keya"])
	fmt.Println(map1)
	map2 := map[string]Point{
		"pointc": Point{1, 3},
		"pointd": Point{2, 5},
	}
	map3 := map[string]Point{
		"pointe": {10, 30},
		"pointf": {20, 50},
	}
	fmt.Println(map2, map3)
	// 插入或修改元素
	map3["pointf"] = Point{1, 3}
	fmt.Println(map3)
	// 获取
	fmt.Println(map3["pointe"])
	// 删除
	delete(map3, "pointe")
	fmt.Println(map3)
	// 双赋值判断元素是否存在
	ele, existFlag := map3["pointfaaaa"]
	fmt.Println(ele, existFlag)
}
func printSlice(name string, slice []int) {
	fmt.Printf("%s %v cap=%d len=%d\n", name, slice, cap(slice), len(slice))
}

type Point struct {
	x, y float64
}
