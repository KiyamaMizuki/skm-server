package server

import (
	"local.com/db-module/model"
	"local.com/db-module/val"
)

func SaveNode(lat float64, lng float64, node_type int, floor int, name string) (node model.Node) {
	db := val.PostgresConnect()
	defer db.Close()
	nodeEx := model.Node{}
	//マイグレーション
	db.AutoMigrate(&nodeEx)
	nodeEx.Latitude = lat
	nodeEx.Longitude = lng
	nodeEx.Type = node_type
	nodeEx.Floor = floor
	nodeEx.Name = name
	db.Create(&nodeEx)
	return nodeEx
}

func NodeGet() (node []model.Node) {
	db := val.PostgresConnect()
	defer db.Close()
	nodeEx := []model.Node{}
	//マイグレーション
	db.AutoMigrate(&nodeEx)
	db.Find(&nodeEx)
	return nodeEx
}
