package test_01

import "fmt"

// 逻辑控制，if/else、for、goto
func init() {

	// if/else
	ifElseFunc()

	// for
	forFunc()

	// goto
	gotoFunc()

	// switch
	switchFunc()

	/*
		// go 函数定义： 可多个参数、多个返回值；
		// 返回值可不声明变量，只声明类型即可，单个返回值也可去掉括号；
		func funcName(input1 type1, input2 type2) (output1 type1, output2 type2) {
			//这里是处理逻辑代码
			// todo

			//返回多个值
			return value1, value2
		}
	*/

	/*
		// 返回值声明变量，return时不需要任何变量
		func SumAndProduct(A, B int) (add int, Multiplied int) {
			add = A+B
			Multiplied = A*B
			return
		}
	*/

	/*
		// 可变长参数
		func myfunc(arg ...int) {}

	*/

}

// if/else
func ifElseFunc() {

	// 计算获取值x,然后根据x返回的大小，判断是否大于10。
	if x := computedValue(3); x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
	//fmt.Println(x)
}

func computedValue(a1 int) int {
	if a1 <= 0 {
		return 0
	}
	return a1 * 5
}

// for
func forFunc() {
	//
	sum1 := 0
	for index := 0; index <= 10; index++ {
		sum1 += index
		// 可内嵌，continue，break
	}
	fmt.Println("sum1 is equal to ", sum1)

	//
	sum2 := 1
	for sum2 <= 10 {
		sum2 += sum2
		// 可内嵌，continue，break
	}
	fmt.Println("sum2 is equal to ", sum2)

	//
	sum3 := 1
	for sum3 <= 10 {
		sum3 += sum3
		// 可内嵌，continue，break
	}
	fmt.Println("sum3 is equal to ", sum3)

	// for 配合 range 读取 map集合的key value
	mp := make(map[string]int)
	mp["a1"] = 1
	mp["a2"] = 2
	mp["a3"] = 3

	for k, v := range mp {
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}

	// 放弃key
	for _, v := range mp {
		//fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}

}

// goto, 其中[hear:] 为标签定义，单词以冒号结束，如[label:]
func gotoFunc() int {
	i := 0
hear:
	println(i)
	if i > 20 {
		return i
	}
	i++
	goto hear
}

// switch, fallthrough 为强制向下执行，哪怕case不满足条件，
// 本函数执行结果为：
// The integer was <= 6
// The integer was <= 7
// The integer was <= 8
// default case
func switchFunc() {

	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
