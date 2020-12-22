/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/7 下午6:09
@Description:
*********************************************/
package model

import (
	"BookOrderSystem/utils"
	"fmt"
)

//User结构体
type User struct {
	Id       int
	UserName string
	PassWord string
	Email    string
}

//CheckUserName 根据用户从数据库中查询一条记录
func (u *User) CheckUserName(userName string) (*User, error) {
	// 写sql语句
	sqlStr := "select id, username,password,email from users where username=?"
	// 执行查询
	row := utils.Db.QueryRow(sqlStr, userName)
	row.Scan(&u.Id, &u.UserName, &u.PassWord, &u.Email)
	return u, nil

}

//CheckUserNameAndPassword 根据用户名和密码从数据库中查询一条记录
func (u *User) CheckUserNameAndPassword(username, password string) (*User, error) {
	// sql
	sqlStr := "select id, username,password, email from users where username= ? and password = ?"
	// 执行查询
	row := utils.Db.QueryRow(sqlStr, username, password)
	err := row.Scan(&u.Id, &u.UserName, &u.PassWord, &u.Email)
	if err != nil {
		fmt.Println(err)
	}

	return u, nil
}

// Insert 向数据库中插入用户信息
// 备注：建表时的字段中间不能有下划线如：usre_name(与orm 框架建表区别）
func (u *User) Insert(username string, password string, email string) error {
	sqlStr := "insert into users (username,password,email) values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
