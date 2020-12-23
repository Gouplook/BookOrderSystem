/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time :  2020/12/7 下午5:36
@Description: 用户基本操作界面

*********************************************/
package controller

import (
	"BookOrderSystem/model"
	"BookOrderSystem/utils"
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
	userModel := new(model.User)
	u, _ := userModel.CheckUserName(username)
	if len(u.Id) > 0 { // 表示用户存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已经存在")
	} else {
		// 获取的数据保存到数据库
		err := userModel.Insert(username, password, email)
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
func Login(w http.ResponseWriter, r *http.Request) {
	// 1.1 判断用户是否已经登陆
	sessModel := model.Session{}
	userModel := model.User{}
	flag, _ := sessModel.IsLogin(r)
	if flag {
		// 跳转图书首页 调用bookhandle GetPageBooksByPrice
		GetPageBooksByPrice(w, r)
	} else {
		// 获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		//调用user model 中验证用户名和密码的方法 CheckUserNameAndPassword
		user, _ := userModel.CheckUserNameAndPassword(username, password)
		// 如果用户名和密码正确 生产UUID作为Session的id
		if len(user.Id) > 0 {
			uuid := utils.CreateUUid()
			//创建一个session
			sess := &model.Session{
				SessionId: uuid,
				UserName:  username,
				UserId:    user.Id,
			}
			//将session数据保存到数据库
			err := sessModel.AddSession(sess)
			if err != nil {
				fmt.Println("session 数据库插入错误 ......")
				return
			}
			// 创建一个cookie ,让其与session相关联
			cookie := http.Cookie{
				Name:     "user", // 不要写username
				Value:    uuid,
				HttpOnly: true,
			}
			// 将cookie发送到浏览器
			http.SetCookie(w, &cookie)
			// 显示注册页面
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "用户名或密码不正确!")
		}
	}
}

//CheckUserName 通过发送Ajax验证用户名是否可用 main调用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	// 1：获取用户输入的用户名
	username := r.PostFormValue("username")
	// 2：调用uesermodel 中验证用户名和密码的方法
	userModel := model.User{}
	user, _ := userModel.CheckUserName(username)
	// 2.1: 用户名存在
	if len(user.Id) > 0 {
		w.Write([]byte("用户名已经存在"))

	} else { // 2.2: 用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))

	}
}

//Logout //处理用户注销的函数 main调用
// 1：获取cookie ，通过cookie，查找session中的用户uid，在数据库对其进行删除
// 2：设置cookie失效，发送到浏览器
func Logout(w http.ResponseWriter, r *http.Request) {
	// 获取Cookie
	cookie, _ := r.Cookie("user")
	sessModel := new(model.Session)
	if cookie != nil {
		// 获取cookie的value值
		cokieValue := cookie.Value
		// 删除sessionId
		sessModel.DeleteSession(cokieValue)
		// 设置cookie 失效
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	// 回到首页
	GetPageBooksByPrice(w, r)

}
