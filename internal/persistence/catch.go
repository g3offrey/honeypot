package persistence

import (
	"github.com/g3offrey/honeypot/pkg/catch"
	"gorm.io/gorm"
)

type CatchModel struct {
	gorm.Model
	catch.Catch
}

func (CatchModel) TableName() string {
	return "catches"
}

type GormCatchRepository struct {
	db *gorm.DB
}

func NewGormCatchRepository(db *gorm.DB) *GormCatchRepository {
	return &GormCatchRepository{db: db}
}

func (r *GormCatchRepository) Create(c catch.Catch) error {
	return r.db.Create(&CatchModel{Catch: c}).Error
}
