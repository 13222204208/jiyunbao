// 自动生成模板AppRelieveStore
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppRelieveStore 结构体
type AppRelieveStore struct {
	global.GVA_MODEL
	Uid      uint   `json:"uid" form:"uid" gorm:"column:uid;comment:申请的用户;"`
	StoreId  uint   `json:"storeId" form:"storeId" gorm:"column:store_id;comment:申请解除合作的门店;"`
	Contents string `json:"contents" form:"contents" gorm:"column:contents;comment:申请解除的原因;size:500;"`
	Status   uint   `json:"status" form:"status" gorm:"column:status;comment:状态;default:0"`
}

// TableName AppRelieveStore 表名
func (AppRelieveStore) TableName() string {
	return "app_relieve_store"
}
