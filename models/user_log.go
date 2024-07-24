package models

import (
	"time"
)

type UserLogs struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	IP         string    `json:"ip"`
	Action     int       `json:"action"`
	Content    string    `json:"content"`
	Expired_at time.Time `json:"expired_at"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
}

const (
	UserLogMobileLogin = 1
)

func CreateLog(log *UserLogs) error {
	return DB.Create(&log).Error
}

func GetLatestValidLogByContent(user_id int, content string) (*UserLogs, error) {
	var log *UserLogs
	err := DB.Where("user_id = ? and content = ? and expired_at > ?", user_id, content, time.Now()).Last(&log).Error
	return log, err
}
