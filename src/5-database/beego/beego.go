package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 导入数据库驱动
	"strconv"
)

const (
	DBHostsIp  = "192.168.0.24:3307"
	DBUserName = "root"
	DBPassWord = "Cobbler1234!"
	DBName     = "test_db"
)

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", DBUserName+":"+DBPassWord+"@tcp("+DBHostsIp+")/"+DBName+"?charset=utf8", 30)
	//注册定义的model
	orm.RegisterModel(new(User))

	// 创建table
	orm.RunSyncdb("default", false, true)

}

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func (h User) String() string {
	return "id: " + strconv.Itoa(h.Id) + "; name: " + h.Name
}

func main() {
	o := orm.NewOrm()

	user := User{Name: "kaywall"}

	// 插入表
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// 更新表
	user.Name = "kaywall_new"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// 读取 one
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Println(u)
	fmt.Printf("ERR: %v\n", err)

	// 删除表
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
