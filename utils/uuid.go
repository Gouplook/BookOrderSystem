/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/9 上午9:22
@Description: 生成32为Uid

*********************************************/
package utils

import (
	"crypto/rand"
	"fmt"
	"log"
)

func CreateUUid()(uuid string){
	u := new([16]byte)
	_,err := rand.Read(u[:])
	if err != nil {
		log.Fatal("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
