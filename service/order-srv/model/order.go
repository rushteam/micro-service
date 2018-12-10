package model

//OrderModel ..
type OrderModel struct {
	// gorm.Model
	SkuID  int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;"`
	Price  int64
	Weight int64
}

//TableName ..
func (OrderModel) TableName() string {
	return "order_order"
}

//GetOrderByNo ...
func (sess *Session) GetOrderByNo(no string) (OrderModel, error) {
	ord := OrderModel{}
	ord.Price = 100 * 30
	ord.SkuID = 1
	ord.Weight = 10
	return ord, nil
}
