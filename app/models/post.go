package models

import "api-backend/sample/app/models/utils"

type Post struct {
	ID       uint `gorm:"primaryKey"`
	Header   string
	Content  *string
	AuthorID int
	Author   User `gorm:"foreignKey:AuthorID"`
	utils.TimestampColumns
}
