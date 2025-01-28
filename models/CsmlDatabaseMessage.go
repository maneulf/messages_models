package models

import "gorm.io/gorm"

type CsmlDataBaseMessageModelDB struct {
	gorm.Model
	Message   string `gorm:"type:text;size:255;not null;" json:"message"`
	UserID    string `gorm:"size:255;not null;" json:"user_id"`
	RequestID string `gorm:"size:255;not null;" json:"request_id"`
	Source    string `gorm:"size:255;not null;" json:"source"`
}
