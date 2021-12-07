package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 错误： := 声明并赋值，同一变量重复赋值
	// a := 1
	// a := 2

	a := 1
	a = 2

	fmt.Println(a)

}

func init() {

	println("go 中四种声明： var 声明变量； const 声明常量； type 声明类； func 声明函数(方法)；")

	println("===>>> go 中四种声明： var 声明变量")

	var tmp1 string = "first temp value"
	var tmp2 = "second temp value"
	tmp3 := "three temp value"

	var ( // 该方式一般用于全局变量
		tmp4 int
		tmp5 float32
	)

	println(tmp1)
	println(tmp2)
	println(tmp3)
	println(tmp4)
	println(tmp5)

	println("===>>> go 中四种声明： const 声明常量，常量的值必须是能够在编译时就能够确定的")

	const p1 = 3.14159       // 隐式声明
	const p2 string = "显示声明" // 显示声明
	println(p1)

	println("===>>> go 中四种声明： type 声明类")

	println("C TO F: ", CToF(AbsoluteZeroC))

	println("F TO C: ", FToC(BoilingC))
}

type celsius float64    // 摄氏度
type fahrenheit float64 // 华氏度

/*
 * 
 */
const (
	AbsoluteZeroC celsius    = -273.15 // 绝对零度
	FreezingC     celsius    = 0       // 结冰点温度
	BoilingC      fahrenheit = 100     // 沸水温度
)

/*
 * c to f
 */
func CToF(c celsius) fahrenheit {
	return fahrenheit(c*9/5 + 32)
}

/*
 * f to c
 */
func FToC(f fahrenheit) celsius {
	return celsius((f - 32) * 5 / 9)
}

func init() {

	fmt.Println("--->>> for i")

	for i := 0; i < 10; i++ {
		fmt.Println("aa")
	}
}

func init() {
	fmt.Println("--->>> for range")

	var arr1 [10]int
	var arr2 [10]int

	for key, value := range arr1 {
		arr2[key] = value
		fmt.Println(key)
	}
	fmt.Println(arr2)

}

func init() {
	println("--->>> iota")

	const (
		i = iota
		j = 3.14
		l
		m    = iota
		k, n = iota, iota
	)
	println(i)
	println(j)
	println(l)
	println(m)
	println(k)
	println(n)

}

func init() {
	println("--->>> uint 可变int型，随操作系统CPU变化，64位操作系统下：int = uint = 8字节， 32位系统下：int = uint = 4字节，int可为负数 uint不可")

	println("--->>> int int8 int16 int32 int64")

	var i1 int = -1 // 可表示负数
	var i2 int8 = 1
	var i3 int16 = 1
	var i4 int32 = 1
	var i5 int64 = 1
	println(unsafe.Sizeof(i1)) // int 占用8字节
	println(unsafe.Sizeof(i2)) // int8 占用1字节
	println(unsafe.Sizeof(i3)) // int16 占用2字节
	println(unsafe.Sizeof(i4)) // int32 占用4字节
	println(unsafe.Sizeof(i5)) // int64 占用8字节

	println("--->>> uint uint8 uint16 uint32 uint64")

	var u1 uint = 1 // 不可表示负数
	var u2 uint8 = 1
	var u3 uint16 = 1
	var u4 uint32 = 1
	var u5 uint64 = 1
	var u6 uintptr = 1
	println(unsafe.Sizeof(u1)) // int 占用8字节
	println(unsafe.Sizeof(u2)) // int8 占用1字节
	println(unsafe.Sizeof(u3)) // int16 占用2字节
	println(unsafe.Sizeof(u4)) // int32 占用4字节
	println(unsafe.Sizeof(u5)) // int64 占用8字节
	println(unsafe.Sizeof(u6))

	println("--->>> pointer uintptr")

	var p1 int = -1
	println(*((*uint)(unsafe.Pointer(&p1))))

	println("--->>> byte == int8 == (java char)")

	var ch1 int8 = 'a'
	var ch2 int16 = '字'
	var ch3 byte = 'a'
	var ch4 byte = 'a'

	println(unsafe.Sizeof(ch1))
	println(unsafe.Sizeof(ch2))
	println(unsafe.Sizeof(ch3))
	println(unsafe.Sizeof(ch4))

}
