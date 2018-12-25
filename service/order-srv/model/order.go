package model

//OrderModel ..
type OrderModel struct {
	// gorm.Model
	OrderNo    string `gorm:"PRIMARY_KEY;"`
	Total      int64
	Discount   int64
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

//GetOrderByOrderNo ...
// func (sess *Session) GetOrderByOrderNo(no string) (OrderModel, error) {
// 	order := OrderModel{}
// 	order.OrderNo = no
// 	return order, nil
// }
