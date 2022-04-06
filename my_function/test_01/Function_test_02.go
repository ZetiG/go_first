package test_01

import (
	"fmt"
	"os"
)

// 逻辑控制，if/else、for、goto、值传递、址传递
func init() {
	println("------------------------>>> Test 02 <<<------------------------")

	// if/else
	ifElseFunc()

	// for
	forFunc()

	// goto
	gotoFunc()

	// switch
	switchFunc()

	// 值传递
	copyVariable()

	// 址传递
	addrVariable()

	// defer
	deferFunc()

	// func 函数作为参数
	funcParamFunc()

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

	// panic
	panicTest()

	// recover
	ret := throwsPanic(panicTest)
	println("===>>>>>> panic,recover: ", ret)

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

// 值传递
func copyVariable() {
	x := 3
	fmt.Println("x = ", x) // 应该输出 "x = 3"

	x1 := copyVariableAdd(x) //调用add1(x)

	fmt.Println("x+1 = ", x1) // 应该输出"x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出"x = 3"
}

// CopyVariableAdd 函数接收的是变量的值的副本
func copyVariableAdd(a int) int {
	a = a + 1
	return a
}

// 址传递
func addrVariable() {
	x := 3

	fmt.Println("x = ", x) // 应该输出 "x = 3"

	x1 := addrVariableAdd(&x) // 调用 add1(&x) 传x的地址

	fmt.Println("x+1 = ", x1) // 应该输出 "x+1 = 4"
	fmt.Println("x = ", x)    // 应该输出 "x = 4"

}

// *int 代表函数接收的是变量的内存地址副本
func addrVariableAdd(a *int) int {
	*a = *a + 1
	return *a
}

// defer 延迟函数，在函数结束退出前调用，多个defer时， 采用后进先出的方式，类似于栈
func deferFunc() {
	a := 1
	defer println("defer after1 =>> ", a) // 这里最后执行，输出3

	a = 2
	defer println("defer after2 =>> ", a) // 这里最后执行，输出3

	a = 3
	defer println("defer after3 =>> ", a) // 这里最后执行，输出3

	a = 4
	println("defer before4 =>> ", a) // 这里输出4

	a = 5
	defer println("defer after5 =>> ", a) // 这里最后执行，输出3

	for i := 0; i < 5; i++ {
		defer println("for defer after i =>> ", i)
	}
}

// 测试func 作为参数使用
func funcParamFunc() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)

	flt := filter(slice, isMoreThenFive) // 函数当做值来传递了
	fmt.Println("more then five slice are: ", flt)
}

type funcParam func(int) bool

// 定义函数，使用func函数作为参数
func filter(arr []int, fp funcParam) []int {
	var result []int

	for _, val := range arr {
		if fp(val) {
			result = append(result, val)
		}
	}
	return result
}

// 定义func函数，作为参数
func isMoreThenFive(a int) bool {
	return a > 3
}

// panic 异常
func panicTest() {
	var user = os.Getenv("USER")

	if user == "" {
		panic("===>>> get env of user is empty")
	}
}

// recover
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}
