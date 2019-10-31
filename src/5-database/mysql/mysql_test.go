package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", DBUserName+":"+DBPassWord+"@tcp("+DBHostsIp+")/"+DBName+"?charset=utf8")
	checkErr(err)
	insert()
	//update(1)
	//query()
	//remove(1)
	//关闭数据库连接
	db.Close()

}
