package model

import "time"

type Admin struct {
	AdminId   int64     `xorm:"pk autoincr"  json:"id"` //主键 自增
	AdminName string    `xorm:"varchar(32)"  json:"admin_name"`
	Status    int       `xorm:"default 0"    json:"status"`
	Avatar    string    `xorm:"varchar(255)" json:"avatar"`
	Password  string    `xorm:"varchar(255)" json:"password"`
	CityName  string    `xorm:"varchar(255)" json:"city_name"`
	CityId    int64     `xorm:"index"        json:"city_id"`
	City      *City     `xorm:"- <- ->"` //所对应的城市结构体（基础表结构体）
	CreatedAt time.Time `xorm:"DateTime"     json:"created_at"`
	UpdatedAt time.Time `xorm:"DateTime"     json:"updated_at"`
}

func (this *Admin) AdminToRespDesc() interface{} {

	respDesc := map[string]interface{}{
		"user_name": this.AdminName,
		"id":        this.AdminId,
		"status":    this.Status,
		"avatar":    this.Avatar,
		"city":      this.CityName,
		"admin":     "管理员",
		"create_at": this.CreatedAt,
	}
	return respDesc
}
