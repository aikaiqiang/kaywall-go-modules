package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const (
	// sqlite3 foo.db # 创建数据库
	// .database # 查看数据库路径
	SQLDBName = "C:/Users/Administrator/foo.db"
)

var db *sql.DB
var err error

func init() {
	fmt.Println("init SQLite DataSource ...")
	db, err = sql.Open("sqlite3", SQLDBName)
	checkErr(err)
}

func main() {
	//insert()
	//update(1)
	remove(1)
	query()
	//关闭数据库连接
	db.Close()
}

func insert() {
	// 准备插入操作
	stmt, err := db.Prepare("INSERT INTO userinfo(username, department, created) values(?,?,?)")
	checkErr(err)
	// 执行插入操作
	res, err := stmt.Exec("kaywall_1", "研发部门", "2019-10-31")
	checkErr(err)
	// 返回最近的自增主键id
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("LastInsertId: ", id)
}

func update(id int) {
	//更新数据
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("kaywall_new", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}

func query() {
	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Printf("uid = [%v], username = [%v], department = [%v], created = [%v]\n", uid, username, department, created)
	}
}

func remove(id int) {
	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
