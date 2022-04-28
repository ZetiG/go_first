package test_01

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func init() {
	println("------------------------>>> Test 03 <<<------------------------")

	// struct 使用
	var per person
	println(per.name)

	httpTest()

}

// 定义类
type person struct {
	name string
	age  int
	sex  byte
}

// 定义接口
type people interface {
}

// 简单的http服务
func httpTest() {
	http.HandleFunc("/", sayHelloName)       //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}
