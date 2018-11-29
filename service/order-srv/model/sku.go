package model

import (
	"errors"
)

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

//GetSkuListBySkuIds ...
func (sess *Session) GetSkuListBySkuIds(skuIds []int64) ([]SkuModel, error) {
	var skuList []SkuModel
	result := sess.Where("sku_id in (?)", skuIds).Scan(&skuList)
	if result.Error != nil {
		return nil, errors.New("没找到商品数据")
	}
	return skuList, nil
}
