package model

type Node struct {
	ID        uint `gorm:"primary_key"`
	Latitude  float64
	Longitude float64
	Floor     int
	Type      int
	Name      string
}
