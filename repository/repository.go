package repository

import (
	"assignment2/structs"

	"gorm.io/gorm"
)

type Repository interface {
	GetOrders() ([]structs.Order, error)
	GetOrderById(id int) (structs.Order, error)
	Create(order structs.Order, item structs.Item) (structs.Order, structs.Item, error)
	Update(order structs.Order) (structs.Order, error)
	Delete(order structs.Order) (structs.Order, error)
}

type Orderepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Orderepository {
	return Orderepository{
		db: db,
	}
}

func (or *Orderepository) GetOrders() ([]structs.Order, error) {
	var orders []structs.Order
	err := or.db.Model(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (or *Orderepository) GetOrderByID(id int) (structs.Order, error) {
	var order structs.Order
	err := or.db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *Orderepository) Create(order structs.Order, item structs.Item) (structs.Order, structs.Item, error) {
	err := r.db.Create(&order, &item).Error
	if err != nil {
		return order, item, err
	}
	return order, item, nil
}

func (r *Orderepository) Update(order structs.Order) (structs.Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (r *Orderepository) Delete(order structs.Order) (structs.Order, error) {
	err := r.db.Delete(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}
