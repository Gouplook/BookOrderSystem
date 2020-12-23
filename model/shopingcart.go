/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/10 下午2:33
@Description:

*********************************************/
package model

import (
	"BookOrderSystem/utils"
)

//Cart 购物车结构体
type Cart struct {
	CartId       string      // 购物车Id
	UserId       int         // 用户Id
	TotalCount   int64       // 购物车所有图书的数目，通过计算得到
	TotalAccount float64     // 购物车图书总金额
	CartItems    []*CartItem // 购物车所有项目
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
	for _, v := range c.CartItems {
		totalAccount += v.GetAmount()
	}
	return totalAccount
}

//AddCart 向购物车表中插入购物车数据
func (c *Cart) AddCart() error {
	//写sql语句
	sqlStr := "insert into carts(id,totalcount,totalamount,userid) values(?,?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, c.CartId, c.GetTotalCount(), c.GetTotalAmount(), c.UserId)
	if err != nil {
		return err
	}
	//获取购物车中的所有购物项
	cartItems := c.CartItems
	//遍历得到每一个购物项
	//cartItemModel := new(CartItem)
	for _, cartItem := range cartItems {
		//将购物项插入到数据库中
		cartItem.AddCartItem()
	}
	return nil
}

//GetCartByUserID 根据用户的id从数据库中查询对应的购物车
func (c *Cart) GetCartByUserID(userId string) (*Cart, error) {
	sqlStr := "select id, totalcount,totalamount,userid from carts where userid = ?"

	row := utils.Db.QueryRow(sqlStr, userId)
	row.Scan(&c.CartId, &c.TotalAccount, &c.TotalCount, &c.UserId)
	//获取当前购物车中所有的购物项
	cartItemModel := new(CartItem)
	cartItems, _ := cartItemModel.GetCartItemsByCartID(c.CartId)
	c.CartItems = cartItems

	return c, nil
}

//UpdateCart 更新购物车中的图书的总数量和总金额
func (c *Cart) UpdateCart() error {
	//sql
	sql := "update carts set totalcount = ? , totalamount = ? where id = ?"
	_,err := utils.Db.Exec(sql,c.GetTotalCount(), c.GetTotalAmount(),c.CartId)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartByCartID 根据购物车的id删除购物车
func (c *Cart) DeleteCartByCartID(cartId string) error {
	// 删除购物车之前需要先删除所有的购物项
	cartItemModel := new(CartItem)
	err := cartItemModel.DeleteCartItemByID(cartId)
	if err != nil {
		return err
	}
	//写sql语句
	sql := "delete from carts where id = ?"
	//执行
	_, err2 := utils.Db.Exec(sql, cartId)
	if err2 != nil {
		return err2
	}
	return nil
}
