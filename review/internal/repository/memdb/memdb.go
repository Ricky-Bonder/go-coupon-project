package memdb

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type Config struct{}

type DbId interface {
	fmt.Stringer
}

type DbEntity interface {
	fmt.Stringer
}

type IRepository[ID DbId, T DbEntity] interface {
	FindByCode(id ID) (T, error)
	Save(e *T) (T, error)
}

//type Repository struct {
//	entries map[string]entity.Coupon
//}

type Repository[ID DbId, T DbEntity] struct {
	db *gorm.DB
}

//func New() *Repository {
//	return &Repository{}
//}

func MakeRepository[ID DbId, T DbEntity](db *gorm.DB) *Repository[ID, T] {
	var t T
	instance := &Repository[ID, T]{}
	instance.db = db
	err := db.AutoMigrate(&t)
	if err != nil {
		return nil
	}
	return instance
}

//func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
//	coupon, ok := r.entries[code]
//	if !ok {
//		return nil, fmt.Errorf("coupon not found")
//	}
//	return &coupon, nil
//}

func (r *Repository[ID, T]) FindByCode(id ID) (T, error) {
	var res T
	err := r.db.Preload(clause.Associations).First(&res, id).Error
	return res, err
}

//func (r *Repository) Save(coupon entity.Coupon) error {
//	r.entries[coupon.Code] = coupon
//	return nil
//}

func (r *Repository[ID, T]) Save(e *T) (T, error) {
	var res T
	var err error
	if err = r.db.Create(e).Error; err != nil {
		log.Println("Error creating " + (*e).String() + " " + err.Error())
		return res, err
	} else {
		log.Println("Created " + (*e).String())
	}
	err = r.db.Find(&res, e).Error
	return res, err
}
