// 自动生成模板MiniUser
package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MiniUser 结构体
type MiniUser struct {
	global.GVA_MODEL
	Nickname      string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:昵称;"`
	Avatar        string `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;"`
	UserNum       string `json:"userNum" form:"userNum" gorm:"column:user_num;comment:联保代码;"`
	MiniType      *int   `json:"miniType" form:"miniType" gorm:"column:mini_type;comment:;"`
	Openid        string `json:"openid" form:"openid" gorm:"column:openid;comment:;"`
	Phone         string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
	RealNameState *int   `json:"realNameState" form:"realNameState" gorm:"column:real_name_state;comment:实名状态;default:0"`
}

// TableName MiniUser 表名
func (MiniUser) TableName() string {
	return "mini_user"
}
