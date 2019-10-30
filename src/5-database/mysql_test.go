package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//const(
//	DBHostsIp  = "192.168.0.24:3307"
//	DBUserName = "root"
//	DBPassWord = "Cobbler1234!"
//	DBName     = "test_db"
//)

func main() {
	db, err := sql.Open("mysql", DBUserName+":"+DBPassWord+"@tcp("+DBHostsIp+")/"+DBName+"?charset=utf8")
	checkErr(err)
	insert(db)
	//update(db, 1)
	//query(db)
	//remove(db, 1)
	//关闭数据库连接
	db.Close()

}
