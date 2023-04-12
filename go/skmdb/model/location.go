package model

type Location struct {
	ID       uint `gorm:"primaryKey"`
	Temp     float64
	Pressure float64
	Created  int64
	Deleted  int64
}
type OpenWeatherMapAPIResponse struct {
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
	Coord   Coord     `json:"coord"`
	Wind    Wind      `json:"wind"`
	Dt      int64     `json:"dt"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
	Pressue  float64 `json:"pressure"`
	Humidity int     `json:"humidity"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}
