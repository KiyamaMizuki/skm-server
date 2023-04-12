package server

import (
	"fmt"
	"math"
	"strconv"

	"local.com/db-module/model"

	_ "github.com/lib/pq"
	"local.com/db-module/val"
)

func SetRoad(nodePointOne int, nodeStart model.Node, nodeEnd model.Node, nodePointTwo int, distance float64, floor int) (road model.Road) {
	db := val.PostgresConnect()
	defer db.Close()

	// fmt.Print("nodeStart -->", nodeStart, "nodeEnd -->", nodeEnd, "\n")
	roadEx := model.Road{
		Distance:    distance,
		Floor:       floor,
		NodeStartID: nodePointOne,
		NodeEndID:   nodePointTwo,
	}
	// fmt.Print("roadEx -->", &roadEx)
	db.Create(&roadEx)
	return roadEx
}

func GetRoad() (road []model.Road) {
	db := val.PostgresConnect()
	defer db.Close()
	roadEx := []model.Road{}
	//マイグレーション
	db.AutoMigrate(&roadEx)
	db.Find(&roadEx)
	return roadEx
}

func CalculateRoute(lat float64, lng float64, floor string, end_node string) (road []model.Result) {
	db := val.GoConnect()
	posgre_db := val.PostgresConnect()
	defer db.Close()
	defer posgre_db.Close()
	nodeList := []model.Node{}
	var start_node string
	posgre_db.Where("floor = ?", floor).Find(&nodeList)
	if len(nodeList) == 0 {
		posgre_db.Where("floor = ?", 1).Find(&nodeList)
	}
	var distance = 100000000.0
	var node_id string
	for i, node := range nodeList {
		new_distance := Distance(lat, lng, node.Latitude, node.Longitude)
		if distance >= new_distance {
			fmt.Println(i)
			distance = new_distance
			node_id = strconv.FormatUint(uint64(node.ID), 10)
		}
	}
	start_node = node_id
	str := "select seq, node, edge, cost from pgr_dijkstra(' SELECT id, node_start_id as source, node_end_id as target, distance as cost from roads', " + start_node + " ," + end_node + ", false)"
	rows, err := db.Query(str)
	if err != nil {
		panic(err.Error())
	}
	resultList := []model.Result{}
	// var results resultList
	// defer rows.Close()
	for rows.Next() {
		var result model.Result
		var node model.Node
		err := rows.Scan(&result.Seq, &result.Node, &result.Edge, &result.Cost)
		nodeRow, _ := db.Query("select * from nodes where id=$1 ", strconv.Itoa(result.Node))
		for nodeRow.Next() {
			err := nodeRow.Scan(&node.ID, &node.Latitude, &node.Longitude, &node.Floor, &node.Type, &node.Name)
			if err != nil {
				panic(err.Error())
			}
		}
		var result_seq = result.Seq
		var result_node = result.Node
		var result_edge = result.Edge
		var result_cost = result.Cost
		var nodeDetail = node
		m := model.Result{
			Seq:        result_seq,
			Node:       result_node,
			Edge:       result_edge,
			Cost:       result_cost,
			NodeDetail: nodeDetail,
		}
		resultList = append(resultList, m)
		if err != nil {
			panic(err.Error())
		}
	}
	return resultList
}
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
