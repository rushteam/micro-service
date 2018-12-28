package model

import "github.com/mlboy/godb/orm"

//SkuModel ..
type SkuModel struct {
	// gorm.Model
	SkuID  int64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;"`
	Price  int64
	Weight int64
}

//TableName ..
func (SkuModel) TableName() string {
	return "product_sku"
}

//GetSkuListBySkuIds ..
func (m *SkuModel) GetSkuListBySkuIds(skuIds []int64) ([]SkuModel, error) {
	var skuList []SkuModel
	orm.Model(m).Where("").FindAll()
	return skuList, nil
}

//SkuModelList ..
type SkuModelList []SkuModel

// //GetSkuListBySkuIds ...
// func (sess *Session) GetSkuListBySkuIds(skuIds []int64) ([]SkuModel, error) {
// 	var skuList []SkuModel
// 	sku := SkuModel{}
// 	sku.Price = 100 * 30
// 	sku.SkuID = 1
// 	sku.Weight = 10
// 	skuList = append(skuList, sku)
// 	return skuList, nil
// 	// result := sess.Where("sku_id in (?)", skuIds).Scan(&skuList)
// 	// if result.Error != nil {
// 	// 	return nil, errors.New("没找到商品数据")
// 	// }
// 	// return skuList, nil
// }
