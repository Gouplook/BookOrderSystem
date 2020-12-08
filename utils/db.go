/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/8 下午4:50
@Description:

*********************************************/
package utils

import "database/sql"
var (
	Db *sql.DB
	err error
)

func init(){
	Db,err = sql.Open("mysql","root:123456@tcp(localhost:3306)/booksystem")
	if err != nil {
		panic(err.Error())
	}
}
