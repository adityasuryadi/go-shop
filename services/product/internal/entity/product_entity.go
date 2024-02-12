package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;" column:"id"`
	Name        string    `gorm:"column:name"`
	Price       int64     `gorm:"column:price"`
	Stock       int64     `gorm:"column:stock"`
	Description string    `gorm:"column:description"`
	CreatedAt   int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt   int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli;default:null"`
}

func (entity *Product) TableName() string {
	return "products"
}

func (entity *Product) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New()
	entity.CreatedAt = time.Now().Unix()
	return nil
}

func (entity *Product) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Unix()
	return nil
}
