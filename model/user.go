/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/7 下午6:09
@Description:
*********************************************/
package model

import (
	"BookOrderSystem/utils"
)
//User结构体
type User struct {
	Id       int
	UserName string
	PassWord string
	Email    string
}

//CheckUserName 根据用户名和密码从数据库中查询一条记录
func (u *User) CheckUserName(userName string) (*User,error) {
	// 写sql语句
	sqlStr := "select id, username,password,email from users where username=?"
	// 执行查询
	row := utils.Db.QueryRow(sqlStr,userName)
	row.Scan(&u.Id,&u.UserName,&u.PassWord,&u.Email)
	return u,nil

}
// Insert 向数据库中插入用户信息
func (u *User)Insert(username string, password string,email string) (error){
	sqlStr := "insert into users(uesrname,password,email) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr,username,password,email)
	if err != nil {
		return err
	}
	return nil
}

//CheckUserNameAndPassword 根据用户名和密码从数据库中查询一条记录
func (u *User)CheckUserNameAndPassword(username ,password string)(*User,error){

	return nil,nil
}
