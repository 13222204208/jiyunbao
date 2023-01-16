// 自动生成模板AppStore
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppStore 结构体
type AppStore struct {
	global.GVA_MODEL
	Uid             uint   `json:"uid" form:"uid" gorm:"column:uid;comment:门店用户;"`
	StoreName       string `json:"storeName" form:"storeName" gorm:"column:store_name;comment:门店名称;"`
	StoreAvatar     string `json:"storeAvatar" form:"storeAvatar" gorm:"column:store_avatar;comment:门店头像;"`
	Category        string `json:"category" form:"category" gorm:"column:category;comment:主营品类;"`
	StorePhone      string `json:"storePhone" form:"storePhone" gorm:"column:store_phone;comment:门店电话;"`
	StoreAddress    string `json:"storeAddress" form:"storeAddress" gorm:"column:store_address;comment:门店地址;"`
	Longitude       string `json:"longitude" form:"longitude" gorm:"column:longitude;comment:门店经度;"`
	Latitude        string `json:"latitude" form:"latitude" gorm:"column:latitude;comment:门店纬度;"`
	Notice          string `json:"notice" form:"notice" gorm:"column:notice;comment:门店公告;"`
	Close           string `json:"close" form:"close" gorm:"column:close;comment:结算模式;"`
	Service         string `json:"service" form:"service" gorm:"column:service;comment:服务设置;"`
	Agreement       string `json:"agreement" form:"agreement" gorm:"column:agreement;comment:合同协议中心;"`
	OperatingStatus int    `json:"operatingStatus" form:"operatingStatus" gorm:"column:operating_status;default:1;comment: 1 营业中，2 暂停，3 满座"`

	Status *int `json:"status" form:"status" gorm:"column:status;comment:合作状态;default:0"`
}

// TableName AppStore 表名
func (AppStore) TableName() string {
	return "app_store"
}
