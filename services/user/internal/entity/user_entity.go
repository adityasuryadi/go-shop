package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;" column:"id"`
	FirtsName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Password  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (entity *User) TableName() string {
	return "users"
}

func (entity *User) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New()
	return nil
}
