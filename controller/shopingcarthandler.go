/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/7 下午5:38
@Description: 购物车操作

*********************************************/
package controller

import (
	"BookOrderSystem/model"
	"fmt"
	"net/http"
)

//AddBook2Cart 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	// 判断用户是否登陆
	sessModel := new(model.Session)
	bookModel := new(model.Book)
	flag, session := sessModel.IsLogin(r)
	if flag {
		// 若登陆获取图书ID
		bookId := r.FormValue("bookId")
		// 根据图书Id获取图书信息
		book, _ := bookModel.GetBookByID(bookId)
		//获取用户id
		userId := session.UserId
		//
		fmt.Println(book)
		fmt.Println(userId)

	}

	// 若登陆获取图书ID

	//获取用户id

	// 判断数据库中是否有当前用户的购物车

}

//GetCartInfo 根据用户的id获取购物车信息

//DeleteCart 清空购物车

//DeleteCartItem 删除购物项

//UpdateCartItem 更新购物项
