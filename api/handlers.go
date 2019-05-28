package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)
/*
 router.POST("/user",CreateUser)
 router.POST("/user/:user_name",Login)
*/

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	 io.WriteString(w,"Create User Handler")
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params){
	username :=p.ByName("user_name")
	io.WriteString(w,username)
}