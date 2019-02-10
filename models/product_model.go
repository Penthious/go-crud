package models

import (
	"go-crud/entities"
)

type ProductModel struct {
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	var products []entities.Product

	// err := productModel.Db.C("product").Find(bson.M{}).All(&products)

	// if err != nil {
	// 	return nil, err
	// }
	return products, nil
}
