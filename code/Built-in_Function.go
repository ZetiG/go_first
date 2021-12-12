package main

import "fmt"

func main() {

	// append
	// appendTest()

	// make
	makeTest()

	// channel、close
	// closeTest()

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
