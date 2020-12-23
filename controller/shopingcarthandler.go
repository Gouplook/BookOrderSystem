/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/7 下午5:38
@Description: 购物车操作

*********************************************/
package controller

import (
	"BookOrderSystem/model"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//AddBook2Cart 添加图书到购物车
func AddBookToCart(w http.ResponseWriter, r *http.Request) {
	// 判断用户是否登陆
	sessModel := new(model.Session)
	bookModel := new(model.Book)
	cartModel := new(model.Cart)
	flag, session := sessModel.IsLogin(r)
	if flag {
		// 若登陆获取图书ID
		bookId := r.FormValue("bookId")
		// 根据图书Id获取图书信息
		book, _ := bookModel.GetBookByID(bookId)
		//获取用户id
		userId := session.UserId
		//判断数据库中是否有当前用户的购物车
		cart, _ := cartModel.GetCartByUserID(userId)
		fmt.Println(book, cart)

	}

	// 若登陆获取图书ID

	//获取用户id

	// 判断数据库中是否有当前用户的购物车

}

//GetCartInfo 根据用户的id获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	sessonModel := new(model.Session)
	_, session := sessonModel.IsLogin(r)
	// 获取用户的Id
	userId := session.UserId
	cartModel := new(model.Cart)
	cart, _ := cartModel.GetCartByUserID(userId)
	if cart != nil {
		//将购物车设置到session中
		session.Cart = cart
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	} else {
		//该用户还没有购物车
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	}

}

//DeleteCart 清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物车的id
	cartId := r.FormValue("cartId")
	//清空购物车
	cartModel := new(model.Cart)
	cartModel.DeleteCartByCartID(cartId)
	GetCartInfo(w, r)

}

//DeleteCartItem 删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request){
	//获取要删除的购物项的id
	cartItemId := r.FormValue("cartItemId")
	//将购物项的id转换为int64
	iCartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	//获取session
	sessionModel := new(model.Session)
	_, session := sessionModel.IsLogin(r)
	//获取用户的id
	userId := session.UserId
	//获取该用户的购物车
	cartModel := new(model.Cart)
	cart, _ := cartModel.GetCartByUserID(userId)
	//获取购物车中的所有的购物项
	cartItems := cart.CartItems
	cartItemModel := new(model.CartItem)
	//遍历得到每一个购物项
	for k, v := range cartItems {
		//寻找要删除的购物项
		if v.CartItemId == iCartItemId {
			//这个就是我们要删除的购物项
			//将当前购物项从切片中移出
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			//将删除购物项之后的切片再次赋给购物车中的切片
			cart.CartItems = cartItems
			//将当前购物项从数据库中删除
			cartItemModel.DeleteCartItemByID(cartItemId)
		}
	}
	//更新购物车中的图书的总数量和总金额
	cartModel.UpdateCart()
	//调用获取购物项信息的函数再次查询购物车信息
	GetCartInfo(w, r)
}

//UpdateCartItem 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request){
	//获取要更新的购物项的id
	cartItemID := r.FormValue("cartItemId")
	//将购物项的id转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取用户输入的图书的数量
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	//获取session
	sessionModel := new(model.Session)
	_, session := sessionModel.IsLogin(r)
	//获取用户的id
	userID := session.UserId
	//获取该用户的购物车
	cartModel := new(model.Cart)
	cart, _ := cartModel.GetCartByUserID(userID)
	//获取购物车中的所有的购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for _, v := range cartItems {
		//寻找要更新的购物项
		if v.CartItemId == iCartItemID {
			//这个就是我们要更新的购物项
			//将当前购物项中的图书的数量设置为用户输入的值
			v.Count = iBookCount
			//更新数据库中该购物项的图书的数量和金额小计
			v.UpdateBookCount()
		}
	}
	//更新购物车中的图书的总数量和总金额
	cartModel.UpdateCart()
	//调用获取购物项信息的函数再次查询购物车信息
	cart, _ = cartModel.GetCartByUserID(userID)
	// GetCartInfo(w, r)
	//获取购物车中图书的总数量
	totalCount := cart.TotalCount
	//获取购物车中图书的总金额
	totalAmount := cart.TotalAccount
	var amount float64
	//获取购物车中更新的购物项中的金额小计
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemID == v.CartItemId {
			//这个就是我们寻找的购物项，此时获取当前购物项中的金额小计
			amount = v.Amount
		}
	}
	//创建Data结构
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	//将data转换为json字符串
	json, _ := json.Marshal(data)
	//响应到浏览器
	w.Write(json)
}
