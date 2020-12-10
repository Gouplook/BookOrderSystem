/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/7 下午5:28

*******************************************/
package main

import (
	"BookOrderSystem/controller"
	"net/http"
)

func main(){

	// 去首页
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//注册
	http.HandleFunc("/regist",controller.Regist)
	//登录
	http.HandleFunc("/login", controller.Login)
	//注销
	http.HandleFunc("logout", controller.Logout)
	// 通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checUserName",controller.CheckUserName)




}
