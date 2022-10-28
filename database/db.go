package database

import (
	"assignment-2/models"
	"errors"

	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func (d Database) CreateOrder(order models.Order) (models.Order, error) {
	dbg := d.db.Create(&order)
	if err := dbg.Error; err != nil {
		return models.Order{}, err
	}

	newOrder := models.Order{
		OrderID:      order.OrderID,
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		Items:        order.Items,
	}

	return newOrder, nil
}

func (d Database) GetOrders() ([]models.Order, error) {
	var orders []models.Order

	err := d.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}

// UpdateOrder return Order result, error, and boolean isFound
func (d Database) UpdateOrder(id int, data models.Order) (models.Order, error, bool) {
	res := d.db.Where("order_id", id).Omit("Items").Updates(&data)

	if res.RowsAffected == 0 && res.Error == nil {
		return models.Order{}, errors.New("order not found"), false
	}

	if res.Error != nil {
		return models.Order{}, res.Error, true
	}

	order := models.Order{OrderID: id}
	for _, v := range data.Items {
		if v.ItemID != 0 {
			d.db.Where("item_id", v.ItemID).Updates(&v)
		}
	}
	err := d.db.Model(&order).Association("Items").Replace(data.Items)

	if err != nil {
		return models.Order{}, err, true
	}

	updatedOrder := models.Order{
		OrderID:      id,
		CustomerName: data.CustomerName,
		OrderedAt:    data.OrderedAt,
		Items:        data.Items,
	}

	return updatedOrder, nil, true
}

// DeleteOrder return Order error, and boolean isFound
func (d Database) DeleteOrder(id int) (error, bool) {
	res := d.db.Delete(&models.Order{}, "order_id", id)

	if res.RowsAffected == 0 && res.Error == nil {
		return errors.New("order not found"), false
	}

	return res.Error, true
}
