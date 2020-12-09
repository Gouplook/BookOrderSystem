/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/9 上午9:48
@Description:

*********************************************/
package model

import (
	"BookOrderSystem/utils"
	"strconv"
)

type Book struct {
	Id     int
	Title  string
	Author string
	Prcie  float64
	Sales  int
	Stock  int 		// 库存
	ImgPth string   // 图书封面
}

//GetBooks 获取数据库中所有的图书

//AddBook 向数据库中添加一本图书

//DeleteBook 根据图书的id从数据库中删除一本图书

//GetBookByID 根据图书的id从数据库中查询出一本图书

//UpdateBook 根据图书的id更新图书信息

//GetPageBooks 获取带分页的图书信息
func (b *Book) GetPageBooks(pageNo string) (*Page, error) {
	//将页码转换为int64类型
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)

	//获取数据库中图书的总记录
	sqlStr := "select count(*) from books"

	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//定义一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	// 获取当前页中的图书
	sqlStr2 := "select id, title,author,prcie,sales,stok,im_pat from book limit ?,?"
	//执行
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*Book
	for rows.Next() {
		book := &Book{}
		rows.Scan(&book.Id,&book.Title,&book.Author,&book.Sales,&book.Stock,&book.ImgPth)
		//将book添加到books中
		books = append(books, book)
	}
	// 创建Page
	page := &Page{
		Books: books,
		PageNo: iPageNo,
		PageSize: pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return  page,nil


}

//GetPageBooksByPrice 获取带分页和价格范围的图书信息
func (b *Book)GetPageBooksByPrice(pageNo string, minPrcie string,maxPrcie string)(*Page, error){
	return nil,nil
}

