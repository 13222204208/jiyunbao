// 自动生成模板MiniClassify
package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MiniClassify 结构体
type MiniClassify struct {
	global.GVA_MODEL
	Title     string `json:"title" form:"title" gorm:"column:title;comment:;"`
	Pid       string `json:"pid" form:"pid" gorm:"column:pid;default:0;comment:;"`
	Icon      string `json:"icon" form:"icon" gorm:"column:icon;comment:;"`
	Level     int    `json:"level" form:"level" gorm:"column:level;default:1;comment:分类等级"`
	AppStatus string `json:"appStatus" form:"appStatus" gorm:"column:app_status;default:1;comment:app商家主营品类显示"`
	Status    string `json:"status" form:"status" gorm:"column:status;default:1;comment:;"`
}

// TableName MiniClassify 表名
func (MiniClassify) TableName() string {
	return "mini_classify"
}
