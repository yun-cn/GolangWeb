package model

type Permission struct {
	PermissionId int64  `xorm:"pk autoincr"    json:"id"`     // 主键id
	Level        string `xorm:"varchar(32)"    json:"level"`  // Level
	Name         string `xorm:"varchar(32)"    json:"name"`   // Name
}