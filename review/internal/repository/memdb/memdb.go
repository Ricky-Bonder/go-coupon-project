package memdb

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

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

type Repository[ID DbId, T DbEntity] struct {
	db *gorm.DB
}

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

func (r *Repository[ID, T]) Get(code string) (T, error) {
	var res T
	err := r.db.Preload(clause.Associations).Clauses(clause.Eq{Column: "code", Value: code}).Find(&res).Error
	return res, err
}

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
