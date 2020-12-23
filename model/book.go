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
	Id     string
	Title  string
	Author string
	Price  float64
	Sales  int
	Stock  int    // 库存
	ImgPth string // 图书封面
}

//GetBooks 获取数据库中所有的图书
func (b *Book)GetBooks()([]*Book, error){
	//sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_pth from books"
	//执行
	rows,_ := utils.Db.Query(sqlStr)

	var books []*Book
	for rows.Next(){
		//book := &Book{}
		rows.Scan(b.Id,b.Title,b.Author,b.Price,b.Sales,b.Stock, b.ImgPth)
		//将b添加到books中
		books = append(books,b)
	}
	return books,nil
}

//AddBook 向数据库中添加一本图书
func (b *Book)AddBook() error{
	// sql 语句
	sqlStr := "insert into books(title,author,price,sales,stock,img_pth) value (?,?,?,?,?,?)"
	// 执行
	_,err := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price, b.Sales, b.Stock, b.ImgPth)
	if err != nil {
		return err
	}
	return nil
}

//DeleteBook 根据图书的id从数据库中删除一本图书
func (b *Book)DeleteBook(bookId string)error{
	// sql语句
	sqlStr := "delete from books where id = ?"
	_,err := utils.Db.Exec(sqlStr,b.Id)
	if err != nil {
		return err
	}
	return nil
}

//GetBookByID 根据图书的id从数据库中查询出一本图书
func (b *Book) GetBookByID(bookId string) (*Book, error) {
	// sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_pth from books where id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookId)
	row.Scan(b.Id, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPth)
	return b, nil
}

//UpdateBook 根据图书的id更新图书信息
func (b *Book)UpdateBook() error {
	// sql语句
	sqlStr := "update books set title=?, author=?, price=?,sales=?,stock=? where id=?"
	// 执行
	_,err := utils.Db.Exec(sqlStr,b.Title, b.Author,b.Price,b.Sales,b.Stock,b.Id)
	if err != nil {
		return err
	}
	return  nil
}

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
	sqlStr2 := "select id, title,author,prcie,sales,stok,img_pth from books limit ?,?"
	//执行
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*Book
	for rows.Next() {
		book := &Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Sales, &book.Stock, &book.ImgPth)
		//将book添加到books中
		books = append(books, book)
	}
	// 创建Page
	page := &Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil

}

//GetPageBooksByPrice 获取带分页和价格范围的图书信息
func (b *Book) GetPageBooksByPrice(pageNo string, minPrice string, maxPrice string) (*Page, error) {
	//将页码转换为int64类型
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sqlStr := "select count(*) from books where price between ? and ?"
	//
	// 查询得到的总记录数
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	var pageSize int64 = 4 // 设置每页显示的条数
	var totalPageSize int64
	if totalRecord%pageSize == 0 {
		totalPageSize = totalRecord / pageSize
	} else {
		totalPageSize = totalRecord/pageSize + 1
	}
	sqlStr2 := "select id,title,author,price,sales,stok,img_pth from books between ? and ? limit ?,?"
	rows, _ := utils.Db.Query(sqlStr2, minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	var books []*Book
	for rows.Next() {
		book := &Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Stock, &book.Sales, &book.ImgPth)
		books = append(books, book)
	}
	// 创建Page
	page := &Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalRecord: totalRecord,
		TotalPageNo: totalPageSize,
	}

	return page, nil
}
