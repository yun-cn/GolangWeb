package model

/**
 * 商店服务类别结构体
 */

type Service struct {
	Id          int     `xorm:"pk autoincr"   json:"id"`         // 主键自增
	Name        string  `xorm:"varchar(32)"   json:"name"`       // 名称
	IconName    string  `xorm:"varchar(255)"  json:"icon_name"`  // Icon name
	IconColor   string  `json:"icon_color"`                      // Icon color
	Description string  `xorm:"varchar(255)" json:"description"` //服务描述
	Shop        []*Shop `xorm:"-"`                               //一个活动可以被多个商店所采用
}
