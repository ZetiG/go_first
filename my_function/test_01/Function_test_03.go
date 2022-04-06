package test_01

func init() {
	println("------------------------>>> Test 03 <<<------------------------")

	// struct 使用
	var per person
	println(per.name)

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
