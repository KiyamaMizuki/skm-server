package model

import "time"

type Authentications struct {
	ID          uint   `gorm:"primaryKey"`
	Mailaddress string `gorm:"size:255"`
	Authcode    int
	CreatedTime time.Time
	DeletedTime time.Time
}
