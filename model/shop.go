package model

type Shop struct {
	Id                          int        `xorm:"pk autoincr"  json:"item_id"`      // 店铺Id
	Name 						string     `xorm:"varchar(32)"  json:"name"`         // Name
	Address                     string     `xorm:"varchar(128)" json:"address"`      // Address
	Latitude                    float32    `json:"latitude"`                         // 经度
	Longitude                   float32    `json:"longitude"`                        // 纬度
	Description                 string     `json:"description"`                      // shop description
	Phone                       string     `json:"phone"`                            // Shop phone
	PromotionInfo				string     `json:"promotion_info"` 					 // Promotion info
	FloatDeliveryFee            int        `json:"float_delivery_fee"`               // 配送费
	FloatMinimumOrderAmount     int        `json:"float_minimum_order_amount"`       // Minimum amount
	IsPremium				    bool       `json:"is_premium"`						 // Is premium
	DeliveryMode		        bool       `json:"delivery_mode"`                    //
	New							bool       `json:"new"`                              // New shop
	Bao                         bool       `json:"bao"`								 // Bao
	Zhun                        bool       `json:"zhun"`			                 // Zhun
	Piao                        bool       `json:"piao"`							 // Piao
	StartTime                   string     `json:"start_time"`                       // Shop start time everyday
	EndTime	                    string     `json:"end_time"`				         // Shop end time everyday
	ImagePath                   string     `json:"image_path"`						 // Image path
	BusinessLicenseImage        string     `json:"business_license_image"`           // Business license image
}
