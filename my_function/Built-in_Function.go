package main

import (
	"fmt"
)

func main() {

	// append
	// appendTest()

	// make
	// makeTest()

	// channel、close
	// closeTest()

	// map
	// mapTest()

	// copy
	// copyTest()

	// complex、real、 imag
	// complexTest()

	// goroutine、panic、recover
	// goroutineTest()
	// println("--->>> panic、recover end <<<---")

	// defer 函数
	// deferA()
	// deferB()
	// i := deferC()
	// fmt.Printf("i: %v\n", i)
	// 	deferD()

	// r1 := deferC1()
	// fmt.Printf("r1: %v\n", r1)

	// r2 := deferC2()
	// fmt.Printf("r2: %v\n", r2)

	// r3 := deferC3()
	// fmt.Printf("r3: %v\n", r3)

	// deferE()

	// panic
	// panicTest()

	panicRecoverTest()
	fmt.Println(" panic recover is run complete")
}

func init() {

	println("append         -- 用来追加元素到数组、slice中,返回修改后的数组、slice	\n" +
		" make           	-- 用来分配内存，返回Type本身(只能应用于slice, map, channel)		\n" +
		" close           	-- 主要用来关闭channel		\n" +
		" delete            -- 从map中删除key对应的value		\n" +
		" panic            	-- 停止常规的goroutine  （panic和recover：用来做错误处理）		\n" +
		" recover         	-- 允许程序定义goroutine的panic动作		\n" +
		" real           	-- 返回complex的实部   （complex、real imag：用于创建和操作复数）		\n" +
		" imag           	-- 返回complex的虚部		\n" +
		" new               -- 用来分配内存，主要用来分配值类型，比如int、struct。返回指向Type的指针		\n" +
		" cap               -- capacity是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）		\n" +
		" copy            	-- 用于复制和连接slice，返回复制的数目		\n" +
		" len               -- 来求长度，比如string、array、slice、map、channel ，返回长度		\n" +
		" print、println     -- 底层打印函数，在部署环境中建议使用 fmt 包		\n")

}

func appendTest() {
	println("---->>> append test...")

	// 定义数组，[3]、[...] 不定长，指定该值不可进行append
	a1 := [...]int{1, 2, 3}
	fmt.Printf("a1: %v\n", a1)

	// append
	str := []string{"h", "e", "l", "l"}
	fmt.Printf("str: %v\n", str)

	str = append(str, "o")
	fmt.Printf("str: %v\n", str)

}

func makeTest() {
	println("---->>> make test...")

	// 对于引用类型的变量，我们不光要声明它，还要为它分配内容空间
	// 对于值类型的声明不需要，是因为已经默认帮我们分配好了
	// 分配内存，Go提供了两种方式，分别是new和make

	// 错误示例
	// var i *int
	// *i = 10
	// fmt.Println(*i)

	// new
	var i *int
	i = new(int)
	*i = 10
	fmt.Println(*i)

	// make 只能应用于slice, map, channel

	// make slice
	// sli1 := make([]int, 5)       // 创建int类型，长度为5的切片
	// sli2 := make([]int, 5, 10)   // 创建int类型，长度为5， 容量为10的切片
	// sli3 := []int{1, 2, 3, 4, 5} // 创建int类型，长度为5 容量为5的切片，如果[3]指定了数值，则代表固定长度的数组而不是切片了

	// make map
	// mp1 := make(map[string]Object)                            		// key 为string，v为Object类型的map
	// mp2 := make(map[string]Object, 5)                         		// key 为string，v为Object类型 初始容量为5的map
	// mp3 := map[string]Object{"key", Object{"aa", "bb", "cc"}} 	// key 为string，v为Object类型 初始容量为5的map

	// make channel
	// ch1 := make(chan int, 5) // 创建int类型，缓存容量为5的通道
	// ch2 := make(chan int)    // 创建int类型，无缓存通道

}

func closeTest() {
	println("---->>> channel close test...")

	// make 创建管道
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go write(ch1)
	go read(ch1, ch2)

	<-ch2

}

/**
 * 向管道ch中循环写入数据
 */
func write(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
	}

	// 关闭管道
	close(ch)
}

/**
 * channel 管道， 操作符 " <- "， 使用 " ch1 <- 1 " 向管道ch1中输入1
 */
func read(ch1 chan int, ch2 chan bool) {

	// for {} 省略条件会无限循环
	for i := 0; i < 10; i++ {
		fmt.Printf("read a int is %d\n", <-ch1)
	}

	ch2 <- true
}

func mapTest() {

	println("---->>> map test...")

	// 定义map
	mp1 := make(map[string]string, 3)

	// 存入数据
	mp1["k1"] = "v1"
	mp1["k2"] = "v2"
	mp1["k3"] = "v3"
	mp1["k3"] = "v4"
	mp1["k5"] = "v5"

	fmt.Printf("mp1: %v\n", mp1)

	// 遍历key
	keys := make([]string, 0, len(mp1))
	for k := range mp1 {
		keys = append(keys, k)
	}
	fmt.Printf("keys: %v\n", keys)

	mp_len := len(mp1)

	fmt.Printf("mp1 length: %v\n", mp_len)

	// range 遍历
	for k, v := range mp1 {
		fmt.Printf("k: %v\n", k)
		fmt.Printf("v: %v\n", v)
	}

	// mp 进行for循环清除
	for k := range mp1 {
		fmt.Printf("循环删除 k: %v\n", k)
		delete(mp1, k)
	}

}

func copyTest() {

	sli1 := []string{"a", "b", "c", "d", "e"}

	fmt.Printf("sli1: %v\n", sli1)

	var sli2 = make([]string, 3)

	fmt.Printf("sli2: %v\n", sli2)

	copy(sli2, sli1)
	println("--->>> cli1 copy to cli2 <<<---")
	fmt.Printf("sli1: %v\n", sli1)
	fmt.Printf("sli2: %v\n", sli2)

}

func complexTest() {
	println("--->>> complex、real、imag <<<---")

	cp := complex(3, 4)

	rl := real(cp)

	fmt.Printf("complex real: %v\n", rl)

	ig := imag(cp)
	fmt.Printf("complex imag: %v\n", ig)

}

func goroutineTest() {

	println("--->>> panic、recover <<<---")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var bar = []int{1}
	fmt.Println(bar[1])

}

// defer 后进先出
func deferD() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
}

// 先入栈，值不随着改变
func deferA() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 后进先出
func deferB() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// defer函数可以读取和重新赋值函数的命名返回参数
func deferC() (i int) {
	defer func() {
		i++
	}()
	return 1
}

func deferC1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func deferC2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func deferC3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func deferE() {
	i := 0
	defer fmt.Println("a:", i)
	//闭包调用，将外部i传到闭包中进行计算，不会改变i的值，如上边的例3
	defer func(i int) {
		fmt.Println("b:", i)
	}(i)
	//闭包调用，捕获同作用域下的i进行计算
	defer func() {
		fmt.Println("c:", i)
	}()
	i++
}

// 调用panic终止线程
func panicTest() {
	panic("crash")
}

func panicRecoverTest() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	panic("crash")

}
