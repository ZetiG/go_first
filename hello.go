package main

import "fmt"

func main() {
	

	// 错误： := 声明并赋值，同一变量重复赋值
	// a := 1
	// a := 2
	
	a := 1
	a = 2

	fmt.Println(a)
}

func init()  {

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
	
	fmt.Println(arr2)

	}
}