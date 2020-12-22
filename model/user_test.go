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


// 用户模块
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
