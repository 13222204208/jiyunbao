// 自动生成模板MiniSet
package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// MiniSet 结构体
type MiniSet struct {
	global.GVA_MODEL
	MiniType     *int   `json:"miniType" form:"miniType" gorm:"column:mini_type;comment:小程序类型;"`
	Appid        string `json:"appid" form:"appid" gorm:"column:appid;comment:;"`
	AppSecret    string `json:"appSecret" form:"appSecret" gorm:"column:app_secret;type:text;comment:秘匙;"`
	MchId        string `json:"mchId" form:"mchId" gorm:"column:mch_id;comment:;"`
	MchSecret    string `json:"mchSecret" form:"mchSecret" gorm:"column:mch_secret;type:text;comment:;"`
	WalletNotice string `json:"walletNotice" form:"walletNotice" gorm:"column:wallet_notice;comment:钱包的公告，个人中心的我的钱包"`
	Public       string `json:"public" form:"public" gorm:"column:public;type:text;comment:;"`
	Private      string `json:"private" form:"private" gorm:"column:private;type:text;comment:;"`
}

// TableName MiniSet 表名
func (MiniSet) TableName() string {
	return "mini_set"
}
