package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func NewDatabase(viper *viper.Viper) *gorm.DB {
	return nil
}
