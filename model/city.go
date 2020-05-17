package model

type City struct {
	CityId    int64   `xorm:"pk autoincr"  json:"id"`        //主键
	CityName  string  `xorm:"varchar(255)" json:"city_name"` // city's Name
	Longitude float32 `xorm:"default 0"    json:"longitude"` //城市经度
	Latitude  float32 `xorm:"default 0"    json:"latitude"`  //城市纬度
	AreaCode  string  `xorm:"varchar(6)"   json:"area_code"` // post code
	Abbr      string  `xorm:"varchar(12)"  json:"abbr"`      //
}
