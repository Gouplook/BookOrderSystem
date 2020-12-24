/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/22 13:48
@Description:

*********************************************/
package model

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}


// 用户模块测试
func TestUser(t *testing.T){
	fmt.Println("测试user中的函数")
	//t.Run("Insert: ",testUser_Insert)
	//t.Run("CheckUserName ：", testUser_CheckUserName)
	t.Run("CheckUserNameAndPassword: ", testUser_CheckUserNameAndPassword)

}

func testUser_Insert(t *testing.T) {

	use := User{
		UserName: "admin",
		PassWord: "123456",
		Email: "admin@163.com",

	}
	err := use.Insert(use.UserName, use.PassWord,use.Email)
	if err != nil{
		fmt.Println(err)
	}
}

func testUser_CheckUserName(t *testing.T) {
	u ,_ := new(User).CheckUserName("jack")
	fmt.Println(u.Id)
}
func testUser_CheckUserNameAndPassword(t *testing.T) {
	u ,_ := new(User).CheckUserNameAndPassword("admin","123456")
	fmt.Println(u.Id)
}

// 图书模块测试
func TestBook(t *testing.T){
	fmt.Println("测试TestBook中的函数")
	//t.Run("GetBooks:", test_GetBooks)
	//t.Run("AddBook:", test_AddBook)
	//t.Run("DeletBook",test_DeleteBook)
	//t.Run("GetBookByID",test_GetBookByID)
	//t.Run("UpdateBook",test_UpdateBook)
	//t.Run("GetPageBooks",test_GetPageBooks)
	t.Run("test_GetPageBooksByPrice",test_GetPageBooksByPrice)
}

func test_GetBooks(t *testing.T){

	b,_ := new(Book).GetBooks()
	for k, v := range b{
		fmt.Printf("%v,%v\n",k, v)
	}

}

func test_AddBook(t *testing.T) {
	book := new(Book)
	book.Title = "王阳明心学2"
	book.Price = 66.5
	book.Author = "王阳明"
	book.Stock = 200
	book.ImgPth = "static/img/default.jpg"
	book.Sales = 50
	book.AddBook()
}

func test_DeleteBook(t *testing.T){
	book := new(Book)
	book.DeleteBook("32")
}

func test_GetBookByID(t *testing.T){

	book, _ := new(Book).GetBookByID("20")
	fmt.Println(book)

}

func test_UpdateBook(t *testing.T){
	book := new(Book)
	book.Title = "王阳明心学223"
	book.Price= 66.5
	book.Author = "王阳明"
	book.Stock = 200
	book.ImgPth= "static/img/default.jpg"
	book.Sales= 50
	book.Id = "31"
	err := book.UpdateBook()
	if err != nil {
		fmt.Println(err)
	}

}
func test_GetPageBooks(t *testing.T) {
	book := new(Book)
	page, _ :=  book.GetPageBooks("6")
	fmt.Println(page)


}
func test_GetPageBooksByPrice(t *testing.T) {
	book := new(Book)
	page, _ := book.GetPageBooksByPrice("5", "10.0", "50.0")
	fmt.Println(page)


}

