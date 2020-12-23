/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/23 16:21
@Description:

*********************************************/
package model

import "BookOrderSystem/utils"

//Order 结构
type Order struct {
	OrderID     string  //订单号
	CreateTime  string  //生成订单的时间
	TotalCount  int64   //订单中图书的总数量
	TotalAmount float64 //订单中图书的总金额
	State       int64   //订单的状态 0 未发货 1 已发货 2 交易完成
	UserID      int64   //订单所属的用户
}


//NoSend 未发货
func (order *Order) NoSend() bool {
	return order.State == 0
}

//SendComplate 已发货
func (order *Order) SendComplate() bool {
	return order.State == 1
}

//Complate 交易完成
func (order *Order) Complate() bool {
	return order.State == 2
}

//AddOrder 向数据库中插入订单
func (order *Order)AddOrder() error{
	//写sql语句
	sql := "insert into orders(id,create_time,totalcount,total_amount,state,user_id) values(?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sql, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}
//GetOrders 获取数据库中所有的订单
func (order *Order)GetOrders()([] *Order, error){
	//写sql语句
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders"
	//执行
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var orders []*Order
	for rows.Next() {
		//order := &model.Order{}
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		orders = append(orders, order)
	}
	return orders, nil
}
//GetMyOrders 获取用户订单
func (order *Order)GetMyOrders(userId string)([]*Order, error){
	//写sql语句
	sql := "select id,create_time,total_count,total_amount,state,user_id from orders where user_id = ?"
	//执行
	rows, err := utils.Db.Query(sql, userId)
	if err != nil {
		return nil, err
	}
	//声明一个切片
	var orders []*Order
	for rows.Next() {
		//创建Order
		//order := &model.Order{}
		//给Order中的字段赋值
		rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		//将Order添加到切片中
		orders = append(orders, order)
	}
	return orders, nil

}
//UpdateOrderState 更新订单的状态，即发货和收货
func (order *Order)UpdateOrderState(orderId string, state int64) error{
	//写sql语句
	sql := "update orders set state = ? where id = ?"
	//执行
	_, err := utils.Db.Exec(sql, state, orderId)
	if err != nil {
		return err
	}
	return nil
}