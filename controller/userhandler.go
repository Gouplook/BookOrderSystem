/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time :  2020/12/7 下午5:36
@Description: 用户基本操作界面
*********************************************/
package controller

import (
	"BookOrderSystem/model"
	"fmt"
	"html/template"
	"net/http"
)

//Regist 处理用户的函注册数
func Regist(w http.ResponseWriter, r *http.Request) {
	// 获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	// 数据查询验证操作
	userModel := model.User{}
	u, _ := userModel.CheckUserName(username)
	if u.Id > 0 { // 表示用户存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已经存在")
	} else {
		// 获取的数据保存到数据库
		err := userModel.Insert(username,password,email)
		if err != nil {
			fmt.Println("Regist Insert failed ...")
			return
		}
		// 用户名称和密码正确
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "")
	}

}

//Login 处理用户登录的函数
func Login(w http.Response, r *http.Request) {
	// 1.1 判断用户是否已经登陆

	//1.2 登陆去显示首页


	// 获取用户名和密码，验证用户名是否存在。


	// 用户名和密码正确，生成UUID作为Session的Id


	// 校验用户名和密码是否正常

	// 响应Login页面
}

//CheckUserName 通过发送Ajax验证用户名是否可用
func CheckUserName() {

}

//Logout //处理用户注销的函数
func Logout(w http.Response, r *http.Request) {
	// 接受注销的用户名

	// 根据用户名，删除

}
