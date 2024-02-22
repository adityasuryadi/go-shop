package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Category struct {
	Id        uuid.UUID             `gorm:"primaryKey;type:uuid;" column:"id"`
	Name      string                `gorm:"column:name"`
	IsActive  int8                  `gorm:"column:is_active"`
	CreatedAt int64                 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64                 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli;default:null"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:milli"`
}

func (entity *Category) TableName() string {
	return "categories"
}

func (entity *Category) BeforeCreate(db *gorm.DB) error {
	// handle create record baru tiap append ke relasy many to many
	if entity.Id == uuid.Nil {
		entity.Id = uuid.New()
	}
	entity.CreatedAt = time.Now().Unix()
	return nil
}

func (entity *Category) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Unix()
	return nil
}
