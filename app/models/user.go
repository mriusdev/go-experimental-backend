package models

import "api-backend/sample/app/models/utils"

type User struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Groups []Group `gorm:"many2many:user_groups"`
	utils.TimestampColumns
}
