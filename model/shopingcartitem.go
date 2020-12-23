/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/10 下午2:36
@Description:

*********************************************/
package model

import "BookOrderSystem/utils"

//CartItem 购物项结构体
type CartItem struct {
	CartItemId int64     // 购物项的Id
	Book       *Book   // 购物项中的图书信息
	Count      int64   // 购物项中图书的数量
	Amount     float64 // 购物项中的图书的金额小计，通过计算得到
	CartId     string  // 当前购物车项属于哪一个购物车
	BookId     string  // **
}

//GetAmount 获取购物项中图书的金额小计，有图书的价格和图书的数量计算得到
func (c *CartItem) GetAmount() float64 {
	// 获取当前购物车中的图书价格
	price := c.Book.Price
	return float64(c.Count) * price
}

//AddCartItem 向购物项表中插入购物项
func (c *CartItem) AddCartItem() error {
	//写sql
	sqlStr := "insert into cartitems(count,amount,bookid,cartid) values(?,?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, c.Count, c.GetAmount(), c.Book.Id, c.CartId)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemByBookIDAndCartID 根据图书的id和购物车的id获取对应的购物项
func (c *CartItem) GetCartItemByBookIDAndCartID(bookId, cartId string) (*CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,cartid from cartitems where bookid = ? and cartid = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookId, cartId)
	//设置一个变量接收图书的id
	//创建cartItem
	err := row.Scan(&c.CartItemId, &c.Count, &c.Amount, &c.CartId)
	if err != nil {
		return nil, err
	}
	//根据图书的id查询图书信息
	bookModel := new(Book)
	book, _ := bookModel.GetBookByID(bookId)
	//将book设置到购物项
	 c.Book = book
	return c, nil
}

//UpdateBookCount 根据购物项中的相关信息更新购物项中图书的数量和金额小计
func (c *CartItem) UpdateBookCount() error{
	//写sql语句
	sql := "update cartitems set count = ? , amount = ? where bookid = ? and cartid = ?"
	//执行
	_, err := utils.Db.Exec(sql, c.Count, c.GetAmount(), c.Book.Id, c.CartId)
	if err != nil {
		return err
	}
	return nil
}

// GetCartItemsByCartID 根据购物车的id获取购物车中所有的购物项
func (c *CartItem) GetCartItemsByCartID(cartId string) ([]*CartItem, error) {
	sqlStr := "select id,count,amount,cartid from cartitem where cartid = ?"
	rows, _ := utils.Db.Query(sqlStr, cartId)
	var cartItems []*CartItem
	for rows.Next() {
		var bookId string
		//创建cartItem
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
func (c *CartItem)DeleteCartItemsByCartID(cartId string) error{
	//写sql语句
	sql := "delete from cartitems where cartid = ?"
	_, err := utils.Db.Exec(sql, cartId)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartItemByID 根据购物项的id删除购物项
func (c *CartItem)DeleteCartItemByID(cartItemId string) error{
	//写sql语句
	sql := "delete from cartitems where id = ?"
	//执行
	_, err := utils.Db.Exec(sql, cartItemId)
	if err != nil {
		return err
	}
	return nil
}
