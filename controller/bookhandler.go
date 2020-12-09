/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/7 下午5:38
@Description: 图书更新或删除基本操作
*********************************************/
package controller

import (
	"BookOrderSystem/model"
	"html/template"
	"net/http"
)

//GetPageBooksByPrice 获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request){
	// 获取页码
	pageNo := r.FormValue("pageNo")
	// 获取价格范围
	minPrice := r.FormValue("min")
	MaxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	bookMode := &model.Book{}
	var page *model.Page
	if minPrice == "" && MaxPrice == "" {
		//调用中获取带分页的图书的函数
		page,_ = bookMode.GetPageBooks(pageNo)

	}else {
		//调用bookmodel中获取带分页和价格范围的图书的函数
		page, _ = bookMode.GetPageBooksByPrice(pageNo,minPrice,MaxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = MaxPrice
	}
	// 查看用户是否已经登陆(session中)
	sessModel := new(model.Session)
	flag, session := sessModel.IsLogin(r)
	if flag {
		//已经登录，设置page中的IsLogin字段和Username的字段值
		page.IsLogin = true
		page.UserName = session.UserName
	}

	//解析模板文件
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)

}



//GetPageBooks 获取带分页的图书



//UpdateOrAddBook 更新或添加图书


//ToUpdateBookPage 去更新或者添加图书的页面



//DeleteBook 删除图书