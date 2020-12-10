/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/10 下午2:33
@Description:

*********************************************/
package model

import (
	"BookOrderSystem/utils"
	"fmt"
)

//Cart 购物车结构体
type Cart struct {
	CartId       int         // 购物车Id
	UserId       int         // 用户Id
	CartItems    []*CartItem // 购物车所有项目
	TotalCount   int64       // 购物车所有图书的数目，通过计算得到
	TotalAccount float64     // 购物车图书总金额
}

//GetTotalCount 获取购物车中图书的总数量
func (c *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range c.CartItems {
		totalCount += v.Count
	}
	return totalCount
}

//GetTotalAmount 获取购物车中图书的总金额
func (c *Cart) GetTotalAmount() float64 {
	var totalAccount float64
	for _, v :=range c.CartItems{
		totalAccount += v.GetAmount()
	}
	return totalAccount
}

//AddCart 向购物车表中插入购物车数据
func (c *Cart)AddCart() error{

	return nil
}


//GetCartByUserID 根据用户的id从数据库中查询对应的购物车
func (c *Cart)GetCartByUserID(userId string) error{
	sqlStr := "select id, total_count,total_amount,user_id from carts where user_id = ?"

	row := utils.Db.QueryRow(sqlStr,userId)
	row.Scan(c.CartId,c.TotalAccount,c.TotalCount,c.UserId)
	//获取当前购物车中所有的购物项
	cartItemModel := new(CartItem)
	//cartItems, _ := cartItemModel.
	fmt.Println(cartItemModel)

	return nil
}

//UpdateCart 更新购物车中的图书的总数量和总金额
func (c *Cart) UpdateCart()error{

	return nil
}

//DeleteCartByCartID 根据购物车的id删除购物车
func (c *Cart) DeleteCartByCartID(cartId string) error{
	return nil
}

