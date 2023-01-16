// 自动生成模板AppAlipayCertification
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppAlipayCertification 结构体
type AppAlipayCertification struct {
	global.GVA_MODEL
	ReqId              string `json:"reqId" form:"reqId" gorm:"column:req_id;comment:;"`
	MercId             string `json:"mercId" form:"mercId" gorm:"column:merc_id;comment:;"`
	MercName           string `json:"mercName" form:"mercName" gorm:"column:merc_name;comment:;"`
	ContactCorpId      string `json:"contactCorpId" form:"contactCorpId" gorm:"column:contact_corp_id;comment:;"`
	ContactsTel        string `json:"contactsTel" form:"contactsTel" gorm:"column:contacts_tel;comment:;"`
	IdHolderType       string `json:"idHolderType" form:"idHolderType" gorm:"column:id_holder_type;comment:;"`
	IdValidDateBegin   string `json:"idValidDateBegin" form:"idValidDateBegin" gorm:"column:id_valid_date_begin;comment:;"`
	IdValidDateEnd     string `json:"idValidDateEnd" form:"idValidDateEnd" gorm:"column:id_valid_date_end;comment:;"`
	BusLincenceBegin   string `json:"busLincenceBegin" form:"busLincenceBegin" gorm:"column:bus_lincence_begin;comment:;"`
	BusLincenceEnd     string `json:"busLincenceEnd" form:"busLincenceEnd" gorm:"column:bus_lincence_end;comment:;"`
	MerContactType     *int   `json:"merContactType" form:"merContactType" gorm:"column:mer_contact_type;comment:;default:65"`
	ManagementType     *int   `json:"managementType" form:"managementType" gorm:"column:management_type;comment:小微经营类型,小微商户必填。 小微经营类型,00-门店场所、01-流动经营/便民服务;"`
	StoreName          string `json:"storeName" form:"storeName" gorm:"column:store_name;comment:门店名称,小微商户必填，根据managementType小微经营类型填写不同值。;"`
	IdTypeCd           string `json:"idTypeCd" form:"idTypeCd" gorm:"column:id_type_cd;comment:;default:0"`
	SendChannelApplyNo string `json:"sendChannelApplyNo" form:"sendChannelApplyNo" gorm:"column:send_channel_apply_no;comment:实名认证申请单单号,该单号用于查询实名认证结果;"`
	QrcodeData         string `json:"qrcodeData" form:"qrcodeData" gorm:"column:qrcode_data;comment:;"`
}

// TableName AppAlipayCertification 表名
func (AppAlipayCertification) TableName() string {
	return "app_alipay_certification"
}
