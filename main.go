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

	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	// 用户模块 ---
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//获取带分页的图书信息
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice",controller.GetPageBooksByPrice)
	//注册
	http.HandleFunc("/regist",controller.Regist)
	//登录
	http.HandleFunc("/login", controller.Login)
	//注销
	http.HandleFunc("logout", controller.Logout)
	// 通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName",controller.CheckUserName)

	http.ListenAndServe(":10086",nil)


	// 购物车模块 ---
	// 添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBookToCart)

	//根据用户的id获取购物车信息
	http.HandleFunc("",controller.GetCartInfo)

}
