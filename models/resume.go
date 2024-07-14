package models

import "time"

type Resumes struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"id"`
	TemplateID int64     `gorm:"column:template_id" json:"template_id"`
	UserID     int64     `gorm:"column:user_id" json:"user_id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
}