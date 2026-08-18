package models

import "api-backend/sample/app/models/utils"

type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Password     string  `gorm:"not null;type:VARCHAR(255)"`
	AccessToken  string  `gorm:"type:TEXT"`
	RefreshToken string  `gorm:"type:TEXT"`
	Groups       []Group `gorm:"many2many:user_groups"`
	utils.TimestampColumns
}

const UserTable string = "users"
