package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


func AddUserCredential(loginNamse string ,pwd string) error{
     stmIns,err :=dbConn.Prepare("INSERT INTO users(login_name,pwd) VALUES(?,?)")
     if(err !=nil){
     	return err
	 }
     _,err=stmIns.Exec(loginName,pwd)
     if err !=nil{
     	returnerr
	 }
     defer stmIns.Close()
     return nil
}

func GetUserCredential(loginNamse string)(string,error) {
	// 准备查询数据
       stmtOut,err :=dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
       if err !=nil{
       	log.Printf("%s",err)
       	return "",err
	   }
       var pwd string
       // login_name写入pwd
       // pwd当参数我们需要改变pwd的值所以用指针
       // 但是当一个变量的改变作用域是在一个作用域下就不需要 比如变量的赋值 pwd=24
       err=stmtOut.QueryRow(loginName).Scan(&pwd)
       if err !=nil&&err!=sql.ErrNoRows{
       	  return "",err
	   }
       defer stmtOut.Close()
       return pwd,nil
}

func DeleteUser(loginName string,pwd string) error{
	// mysql是以列为查找 所以删除用户我们需要删除其用户名和密码（还有其他列 全部的列用login_name和pwd去定位---类似主键	）
	stmtDel,err:=dbConn.Prepare("DELETE * FROM users WHERE login_name=? AND pwd=?")
	if err !=nil{
		log.Printf("DELETEUser error")
		return err
	}
	_,err=stmtDel.Exec(loginName,pwd)
	if err!=nil{
		return err
	}
	stmtDel.Close()
	return nil
}

// Exec 和 QueryROw都是执行sql语句
// 所以学习go直接去看它库的方法说明就行 --它就是文档
// 前端也是 直接定位其方法 看解释再去搜文档 要相信自己的理解能力