package models

import "api-backend/sample/app/models/utils"

type Group struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Permissions []Permission `gorm:"many2many:group_permissions"`
	utils.TimestampColumns
}
