package service

import (
	postgres "distributorBE/database"
	"distributorBE/model"

	"github.com/jinzhu/gorm"
)

// BrandService brand service CRUD
type BrandService struct {
	model.Brand
}

var db *gorm.DB

func init() {
	db = postgres.GetDbConnection()
	db = db.LogMode(true)
}

// FindByID ...
func (brandSerice *BrandService) FindByID(id int64) model.Brand {

	brand := model.Brand{}
	db.Where("id = ? ", id).Find(&brand)

	return brand
}
