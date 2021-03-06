/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/9 上午9:19
@Description:

*********************************************/
package model

import (
	"BookOrderSystem/utils"
	"net/http"
)

//Session 结构
type Session struct {
	SessionId string
	UserName  string
	UserId    string
	Cart *Cart
	OrderId string
	//Orders []*
}

//AddSession 向数据库中添加Session
func (s *Session) AddSession(sess *Session) (err error) {
	// sql语句
	sqlStr := "insert into session values(?,?,?)"
	// 执行插入
	_, err = utils.Db.Exec(sqlStr,sess.SessionId,sess.UserId,sess.UserName)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 删除数据库中的Session
func (s *Session)DeleteSession(sessId string )(err error){
	// sql 语句
	sqlStr := "delete from sessions where session_id = ?"

	// 执行
	_, err = utils.Db.Exec(sqlStr,sessId)
	if err != nil {
		return err
	}
	return nil
}

//GetSession 根据session的Id值从数据库中查询Session
func (s *Session) GetSession(sessId string) (*Session, error) {
	//写sql语句
	sqlStr := "select session_id,username,user_id from session where session_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessId)
	// 创建Session
	sess := &Session{}
	row.Scan(&sess.SessionId, &sess.UserName, &sess.UserId)
	return sess,nil
}
//IsLogin 判断用户是否已经登录 false 没有登录 true 已经登录
func (s *Session) IsLogin(r *http.Request)(bool, *Session){
	//根据Cookie的name获取Cookie
	cookie,_ := r.Cookie("user")
	// 没有登陆，登陆后添加cookie
	if cookie != nil{
		//获取Cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库中查询与之对应的Session
		session,_ := s.GetSession(cookieValue)
		if len(session.UserId) >0{
			//已经登陆
			return true,session
		}
	}
	// 没有登陆
	return false,nil
}

