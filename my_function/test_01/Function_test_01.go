package test_01

import (
	"fmt"
	"os"
)

func init() {
	println("------------------------>>> Test 01 <<<------------------------")

	Test01()
}

func Test01() {
	println("Test_01 ==========>>>>>>>>>>")

	// 此处只是声明，无法add, 但是可以获取值，返回空
	var mp map[string]string
	println(mp["a"])

	// 此处为声明并初始化一个map，方式1
	var mp1 = map[string]string{"a": "a1"}

	mp1["b"] = "a2"
	println(mp1["a"])
	println(mp1["b"])

	// 此处为声明并初始化一个map，方式2, 可知道容量
	//mp2 := make(map[string]string)
	mp2 := make(map[string]string, 2)

	mp2["b1"] = "b11"
	println(mp2["b1"])
	mp2["b2"] = "b22"
	println(mp2["b2"])
	mp2["b3"] = "b33"
	println(mp2["b3"])
	mp2["b4"] = "b44"
	println(mp2["b4"])

	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	// 当存在该值时，value返回该key对应的value， 且ok=true, 否则value=0 ok=false
	value, ok := rating["C"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", value)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	delete(rating, "C") // 删除key为C的元素

	buf := make([]byte, 1024)
	//f, _ := os.Open("/my_function/test_01/test.text")
	//f, _ := os.Open("/Users/zhangmengke/go/go_first/my_function/test_01/test.text")
	f, err := os.Open("/Users/zhangmengke/go/go_first/my_function/test_01/test.text")
	fmt.Printf("err: %v\n", err)

	defer f.Close()

	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break

		}
		os.Stdout.Write(buf[:n])
	}

}
