package config

import "github.com/jinzhu/gorm"

type Env struct {
	DB *gorm.DB
}
