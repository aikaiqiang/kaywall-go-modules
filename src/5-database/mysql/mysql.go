package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHostsIp  = "192.168.0.24:3307"
	DBUserName = "root"
	DBPassWord = "Cobbler1234!"
	DBName     = "test_db"
)

var db *sql.DB
var err error

func init() {
	fmt.Println("init DataSource ...")
	db, err = sql.Open("mysql", DBUserName+":"+DBPassWord+"@tcp("+DBHostsIp+")/"+DBName+"?charset=utf8")
	checkErr(err)
}

func main() {
	//insert(db)
	//update(db, 1)
	//remove(db, 1)
	query()
	//关闭数据库连接
	db.Close()

}

func insert() {
	// 准备插入操作
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)
	// 执行插入操作
	res, err := stmt.Exec("kaywall", "研发部门", "2019-10-31")
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

	res, err := stmt.Exec("aikq", id)
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
