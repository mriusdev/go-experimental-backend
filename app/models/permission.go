package models

import "api-backend/sample/app/models/utils"

type Permission struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	utils.TimestampColumns
}
