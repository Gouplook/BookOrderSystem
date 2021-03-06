/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/7 下午5:38
@Description: 图书更新或删除基本操作

*********************************************/
package controller

import (
	"BookOrderSystem/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//GetPageBooksByPrice 获取带分页和价格范围的图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
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
	// 没有输入价格
	if minPrice == "" && MaxPrice == "" {
		//调用中获取带分页的图书的函数
		page, _ = bookMode.GetPageBooks(pageNo)

	} else {
		//调用bookmodel中获取带分页和价格范围的图书的函数
		page, _ = bookMode.GetPageBooksByPrice(pageNo, minPrice, MaxPrice)
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
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	// 获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	// 调用model中book 获取带分页的图书函数
	bookModel := new(model.Book)
	page, _ := bookModel.GetPageBooks(pageNo)
	// 解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, page)
}

//UpdateOrAddBook 更新或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	// 获取图书基本信息
	bookId := r.FormValue("bookId")
	title := r.FormValue("title")
	author := r.FormValue("author")
	price := r.FormValue("price")
	sales := r.FormValue("sales")
	stock := r.FormValue("stock")

	// 插入到数据库中的数据需要转换
	fPrice, _ := strconv.ParseFloat(price, 64)
	//iBookId, _ := strconv.Atoi(bookId)
	iSales, _ := strconv.Atoi(sales)
	iStock, _ := strconv.Atoi(stock)
	book := &model.Book{
		Id: bookId,
		Title: title,
		Author: author,
		Price: fPrice,
		Sales: iSales,
		Stock: iStock,
		ImgPth: "/static/img/default.jpg",

	}
	if len(book.Id) > 0 {
		// 更新
		err := book.UpdateBook()
		if err != nil {
			fmt.Println("book update failed......")
			return
		}
	} else {
		// 添加
		err := book.AddBook()
		if err != nil {
			fmt.Println("book add failed......")
			return
		}
	}

}

//ToUpdateBookPage 去更新或者添加图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	// 获取需要更新图书的Id
	bookId := r.FormValue("bookId")

	// 调用model中book获取图书的函数
	bookModel := new(model.Book)
	book, err := bookModel.GetBookByID(bookId)
	if err != nil {
		fmt.Println("更新图书失败......")
		return
	}
	if len(book.Id) > 0 {
		//在更新图书
		//解析模板 (编辑页）
		t := template.Must(template.ParseFiles("views/pages/manger/book_edit.html"))
		// 执行
		t.Execute(w, book)
	} else {
		//在添加图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manger/book_edit.html"))
		t.Execute(w, "")
	}
}

//DeleteBook 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	bookModel := new(model.Book)
	err := bookModel.DeleteBook(bookId)
	if err != nil {
		return
	}
	// 回到 获取带分页的图书 再次查询一次数据库（***）
	GetPageBooks(w, r)
}
