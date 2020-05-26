package model

//OrderModel ..
type OrderModel struct {
	OrderNo    string `db:"order_no,pk"`
	Total      int64  `db:"order_no,pk"`
	Discount   int64  `db:"order_no,pk"`
	Payment    int64
	PayState   int64
	OrderState int64
	// string pay_at =7;
	// //地址
	// Address address=9;
	// //更新时间
	// string updated_at = 10;
	// //创建时间
	// string created_at = 11;
}

//TableName ..
func (OrderModel) TableName() string {
	return "order_order"
}

//GetOrderByOrderNo ..
func (m *OrderModel) GetOrderByOrderNo(no string) error {
	order := OrderModel{}
	order.OrderNo = no
	gosql.Fetch(order)
	return nil
}