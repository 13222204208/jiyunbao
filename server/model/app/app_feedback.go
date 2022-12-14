// 自动生成模板AppFeedback
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppFeedback 结构体
type AppFeedback struct {
	global.GVA_MODEL
	Uid      uint   `json:"uid" form:"uid" gorm:"column:uid;comment:反馈的用户ID;size:10;"`
	Contents string `json:"contents" form:"contents" gorm:"column:contents;comment:反馈内容;"`
}

// TableName AppFeedback 表名
func (AppFeedback) TableName() string {
	return "app_feedback"
}
