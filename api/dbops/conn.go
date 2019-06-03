package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// 为什么变量类型是指针类型 有什么作用说明这个变量是sql.DB的指针类型dbConn就是个指针
	// 其实指针的目的就是为了可以修改那个值而不是拷贝 在go语言中
	//在这个包同样拿到这个变量
	dbConn *sql.DB
	err error
)

func init(){
	dbConn,err :=sql.Open("mysql","root:123123lxs@#@tcp(localhost:3306)/video_server?charset=utf8")
	if err !=nil{
		panic(err.Error())
	}
}