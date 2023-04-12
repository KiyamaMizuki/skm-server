package skmdb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"local.com/db-module/model"
	"local.com/db-module/val"
)

// Gene メールを格納し、コードを返す
func CalculateAltitude(currentPressure string) float64 {
	db := val.MysqlConnect()
	defer db.Close()
	err := godotenv.Load(fmt.Sprintf(".env"))
	if err != nil {
	}
	api := os.Getenv("Weather")
	db.AutoMigrate(&model.Location{})
	current, _ := strconv.ParseFloat(currentPressure, 64)
	i := model.Location{}
	recodeNotFound := db.Table("locations").Where("deleted >= ?", time.Now().Unix()).First(&i).RecordNotFound()
	if recodeNotFound {
		token := api                                                  // APIトークン
		endPoint := "https://api.openweathermap.org/data/2.5/weather" // APIのエンドポイント
		values := url.Values{}
		values.Set("lat", "26.286520")
		values.Set("lon", "127.739426")
		values.Set("APPID", token)
		res, err := http.Get(endPoint + "?" + values.Encode())
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		// レスポンスを読み取り
		bytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
		var apiRes model.OpenWeatherMapAPIResponse
		if err := json.Unmarshal(bytes, &apiRes); err != nil {
			panic(err)
		}
		location := model.Location{}
		location.Pressure = apiRes.Main.Pressue
		location.Temp = apiRes.Main.Temp
		location.Created = time.Now().Unix()
		location.Deleted = time.Now().Add(10 * time.Minute).Unix()
		if db.Model(&location).Where("id = ?", 1).Updates(&location).RowsAffected == 0 {
			db.Create(&location)
		}
		print("waashi")
		h := (math.Pow((location.Pressure/current), 1/5.257) - 1) * (location.Temp) / 0.0065
		return h
	} else {
		print("waasd")
		db.Table("locations").First(&i)
		h := (math.Pow((i.Pressure/current), 1/5.257) - 1) * (i.Temp) / 0.0065
		return h
	}
}
