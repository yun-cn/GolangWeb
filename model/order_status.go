package model

type OrderStatus struct {
	Id         int    `xorm:"pk antoincr"   json:"id"`
	StatusId   int    `json:"status_id"`
	StatusDesc string `xorm:"varchar(255)"`
}
