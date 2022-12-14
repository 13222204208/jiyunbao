// 自动生成模板AppWechatCertification
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppWechatCertification 结构体
type AppWechatCertification struct {
	global.GVA_MODEL
	ReqId                  string `json:"reqId" form:"reqId" gorm:"column:req_id;comment:;"`
	MercId                 string `json:"mercId" form:"mercId" gorm:"column:merc_id;comment:;"`
	MercName               string `json:"mercName" form:"mercName" gorm:"column:merc_name;comment:;"`
	IdTypeCd               string `json:"idTypeCd" form:"idTypeCd" gorm:"column:id_type_cd;comment:联系人证件类型;default:00"`
	IdFrontImg             string `json:"idFrontImg" form:"idFrontImg" gorm:"column:id_front_img;comment:;"`
	IdBackImg              string `json:"idBackImg" form:"idBackImg" gorm:"column:id_back_img;comment:;"`
	ContIdValidDateBegin   string `json:"contIdValidDateBegin" form:"contIdValidDateBegin" gorm:"column:cont_id_valid_date_begin;comment:;"`
	ContIdValidDateEnd     string `json:"contIdValidDateEnd" form:"contIdValidDateEnd" gorm:"column:cont_id_valid_date_end;comment:;"`
	ContactCorpId          string `json:"contactCorpId" form:"contactCorpId" gorm:"column:contact_corp_id;comment:;"`
	IdAddress              string `json:"idAddress" form:"idAddress" gorm:"column:id_address;comment:企业商户必填;"`
	MerContactType         *int   `json:"merContactType" form:"merContactType" gorm:"column:mer_contact_type;comment:联系人类型;default:65"`
	ManagementType         *int   `json:"managementType" form:"managementType" gorm:"column:management_type;comment:小微经营类型,小微商户必填。 小微经营类型,00-门店场所、01-流动经营/便民服务"`
	StoreName              string `json:"storeName" form:"storeName" gorm:"column:store_name;comment:门店名称,小微商户必填，根据managementType小微经营类型填写不同值。;"`
	BusAutLetterImg        string `json:"busAutLetterImg" form:"busAutLetterImg" gorm:"column:bus_aut_letter_img;comment:业务办理授权函,联系人类型为经办人时必填;"`
	SfzfrontPhoto          string `json:"sfzfrontPhoto" form:"sfzfrontPhoto" gorm:"column:sfzfront_photo;comment:法人身份证正面照,图片base64编码;type:text; "`
	SfzbackPhoto           string `json:"sfzbackPhoto" form:"sfzbackPhoto" gorm:"column:sfzback_photo;comment:法人身份证反面照,图片base64编码，图片大小最大2M;type:text;"`
	ShopPhoto              string `json:"shopPhoto" form:"shopPhoto" gorm:"column:shop_photo;comment:门头照;type:text;"`
	StoreEnvirPhoto        string `json:"storeEnvirPhoto" form:"storeEnvirPhoto" gorm:"column:store_envir_photo;comment:店内环境照片,图片base64编码，报备时未上送则必填;type:text;"`
	ApplyNo                string `json:"applyNo" form:"applyNo" gorm:"column:apply_no;comment:申请单编号"`
	SendChannelAuthApplyId string `json:"sendChannelAuthApplyId" form:"sendChannelAuthApplyId" gorm:"column:send_channel_auth_apply_id;comment:实名认证申请单编号,银盛系统生成的申请单编号"`
	QrcodeData             string `json:"qrcodeData" form:"qrcodeData" gorm:"column:qrcode_data;comment:小程序码图片,base64编码;type:text"`
	ThirdMercIdX           string `json:"thirdMercIdX" form:"thirdMercIdX" gorm:"column:third_merc_id_x;comment:第三方商户号"`
	AuthStateX             string `json:"authStateX" form:"authStateX" gorm:"column:auth_state_x;comment:授权状态"`
}

// TableName AppWechatCertification 表名
func (AppWechatCertification) TableName() string {
	return "app_wechat_certification"
}
