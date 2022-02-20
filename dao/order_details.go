/*******
* @Author:qingmeng
* @Description:
* @File:order_details
* @Date2022/2/19
 */

package dao

import (
	"shop/model"
	"time"
)

type OrderDao struct {
}

// CreateOrder 传入addressId生成订单
func (d *OrderDao) CreateOrder(addressId int) (int, error) {
	rs, err := DB.Exec("insert into shop.order_details (address_id, date) VALUES(?,?) ", addressId, time.Now())
	id, err := rs.LastInsertId()
	return int(id), err
}

// UpdateOrder 更新订单
func (d *OrderDao) UpdateOrder(order model.OrderDetails) error {
	_, err := DB.Exec("update shop.order_details set address_id=?,order_state=?,confirm_receipt=?,money=? where order_id=?", order.AddressId, order.OrderState, order.ConfirmReceipt, order.Money, order.OrderId)
	return err
}

// SelectOrderByOrderId 根据oderId返回订单
func (d *OrderDao) SelectOrderByOrderId(orderId int) (model.OrderDetails, error) {
	var order model.OrderDetails
	row := DB.QueryRow("select * from shop.order_details where order_id=? ", orderId)
	if row.Err() != nil {
		return order, row.Err()
	}
	err := row.Scan(&order.OrderId, &order.AddressId, &order.Date, &order.OrderState, &order.ConfirmReceipt, &order.Money)
	return order, err
}

func (d *OrderDao) SelectOrdersByUid(uid int) ([]model.OrderDetails, error) {
	var orders []model.OrderDetails
	rows, err := DB.Query("select order_details.order_id, address_id, date, order_state, confirm_receipt, money from shop.shopping_cart,shop.order_details where order_details.order_id=shopping_cart.order_id and uid=?", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order model.OrderDetails
		err = rows.Scan(&order.OrderId, &order.AddressId, &order.Date, &order.OrderState, &order.ConfirmReceipt, &order.Money)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, err
}
