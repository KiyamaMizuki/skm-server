package model

import "time"

type Users struct {
	ID          uint   `gorm:"primaryKey"`
	Mailaddress string `gorm:"size:255"`
	Token       string //`json:"pass"`
	Username    string
	CreatedTIME time.Time
	DeletedTIME time.Time
}
