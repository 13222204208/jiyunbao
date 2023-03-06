// 自动生成模板AppUser
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mini"
)

// AppUser 结构体
type AppUser struct {
	global.GVA_MODEL
	Nickname      string            `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;"`
	Password      string            `json:"password" form:"password" gorm:"column:password"`
	Phone         string            `json:"phone" form:"phone" gorm:"column:phone;comment:;size:11;"`
	UserNum       string            `json:"userNum" form:"userNum" gorm:"column:user_num;comment:用户编号;"`
	ServiceId     *int              `json:"serviceId" form:"serviceId" gorm:"column:service_id;comment:服务机构id;"`
	JoinCode      string            `json:"joinCode" form:"joinCode" gorm:"column:join_code;comment:联保代码;"`
	MchType       *int              `json:"mchType" form:"mchType" gorm:"column:mch_type;comment:商户类型 1联盟，2传统;default:2"`
	Agreement     string            `json:"agreement" form:"agreement" gorm:"column:agreement;comment:商户入驻协议;"`
	CategoryId    int               `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:主营品类ID"`
	Classify      mini.MiniClassify `gorm:"foreignKey:CategoryId"`
	AllianceRatio float64           `json:"allianceRatio" form:"allianceRatio" gorm:"column:alliance_ratio;type:decimal(4,2);default:10;comment:联盟商户服务费比例"`
	OrdinaryRatio float64           `json:"ordinaryRatio" format:"ordinaryRatio" gorm:"column:ordinary_ratio;type:decimal(4,2);default:0.38;comment:一般商户收单手续费"`
	MchState      int               `json:"mchState" form:"mchState" gorm:"column:mch_state;comment:商户认证的状态 0未认证 1 审核中 ，2 认证成功，3 认证未通过 ;default:0"`
	PayState      int               `json:"payState" form:"payState" gorm:"column:pay_state;default:0;comment:支付认证状态 0未认证 1 审核中 ，2 认证成功，3 认证未通过"`
	IsInstitution int               `json:"isInstitution" form:"isInstitution" gorm:"is_institution;comment:是否成为机构;default:0"`
	Avatar        string            `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;"`
	Wechat        string            `json:"wechat" form:"wechat" gorm:"column:wechat;comment:;"`
	Code          string            `json:"code" form:"code" gorm:"-"`
}

// TableName AppUser 表名
func (AppUser) TableName() string {
	return "app_user"
}
