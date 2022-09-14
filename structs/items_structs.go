package structs

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Item_Id     uint `gorm:"primaryKey"`
	Item_Code   uint
	Description string
	Quantity    uint
	Order_Id    uint
}
