package model

type Registrations struct {
	ID        uint `gorm:"primaryKey"`
	Userid    int
	Username  string
	Classid   int
	Classname string
}
