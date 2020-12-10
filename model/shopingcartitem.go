/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/10 下午2:36
@Description:

*********************************************/
package model

import "BookOrderSystem/utils"

//CartItem 购物项结构体
type CartItem struct {
	CartItemId int     // 购物项的Id
	Book       *Book   // 购物项中的图书信息
	Count      int64   // 购物项中图书的数量
	Amount     float64 // 购物项中的图书的金额小计，通过计算得到
	CartId     string  // 当前购物车项属于哪一个购物车
}

//GetAmount 获取购物项中图书的金额小计，有图书的价格和图书的数量计算得到
func (c *CartItem) GetAmount() float64 {
	// 获取当前购物车中的图书价格
	price := c.Book.Price
	return float64(c.Count) * price
}

//AddCartItem 向购物项表中插入购物项

//GetCartItemByBookIDAndCartID 根据图书的id和购物车的id获取对应的购物项

//UpdateBookCount 根据购物项中的相关信息更新购物项中图书的数量和金额小计

// GetCartItemsByCartID 根据购物车的id获取购物车中所有的购物项
func (c *CartItem) GetCartItemsByCartID(cartId string) ([]*CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cartitem where cart_id = ?"
	rows, _ := utils.Db.Query(sqlStr, cartId)
	var cartItems []*CartItem
	for rows.Next() {
		var bookId string
		cartItem := &CartItem{}
		rows.Scan(&cartItem.CartId, &cartItem.Count, &cartItem.Amount, &bookId, &cartItem.CartItemId)
		// 根据bookId获取图书信息
		bookModel := new(Book)
		book, _ := bookModel.GetBookByID(bookId)
		// 将book设置到购物项中
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil

}

//DeleteCartItemsByCartID 根据购物车的id删除所有的购物项

//DeleteCartItemByID 根据购物项的id删除购物项
