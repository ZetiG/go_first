package test_01

import (
	"database/sql"
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

	//创建表
	sql_table := `
    CREATE TABLE IF NOT EXISTS t_user(
        id INT(10) PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(64) DEFAULT NULL,
        age INT(2) DEFAULT NULL
    );`
	open.Exec(sql_table)

	ins := "INSERT INTO t_user(name, age) values(?, ?)"
	stmt, err := open.Prepare(ins)
	checkErr(err)

	exec, err := stmt.Exec("zhangsan", 18)
	checkErr(err)

	lastInsId, err := exec.LastInsertId()
	checkErr(err)

	println(lastInsId)

}

/*
	异常输出
*/
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
