package structs

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Item_Id     uint   `gorm:"column:id"`
	Item_Code   uint   `gorm:"column:item_code" json:"item_code"`
	Description string `gorm:"column:description" json:"description"`
	Quantity    uint   `gorm:"column:quantity" json:"quantity"`
	Order_Id    uint   `gorm:"column:order_id" json:"id"`
}
