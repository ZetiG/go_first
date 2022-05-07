package test_01

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

// 连接数据库
func init() {
	println("------------------------>>> Test 04 <<<------------------------")

	registerDbDriverTest()
}

const (
	dbDriverName = "sqlite3"
	dbName       = "/Users/zhangmengke/go/go_first/my_function/test_01/sqlite3_test.db"
)

func registerDbDriverTest() {

	open, err := sql.Open(dbDriverName, dbName)
	checkErr(err)

	// insert
	ins := "INSERT INTO t_user(name, age) values(?, ?)"
	stmt, err := open.Prepare(ins)
	checkErr(err)

	exec, err := stmt.Exec("lisi", 20)
	checkErr(err)

	lastInsId, err := exec.LastInsertId()
	checkErr(err)

	println(lastInsId)

	// select
	rows, err := open.Query("SELECT * FROM t_user")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		var age int

		err = rows.Scan(&id, &name, &age)
		checkErr(err)

		fmt.Print(id)
		fmt.Print(name)
		fmt.Println(age)

	}

}

/*
	异常输出
*/
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
