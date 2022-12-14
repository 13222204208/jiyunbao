// 自动生成模板AppPay
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"

)

// AppPay 结构体
type AppPay struct {
      global.GVA_MODEL
      Uid  uint `json:"uid" form:"uid" gorm:"column:uid;comment:;"`
      BusinessLicense  string `json:"businessLicense" form:"businessLicense" gorm:"column:business_license;comment:营业执照原件;"`
      OpenCard  string `json:"openCard" form:"openCard" gorm:"column:open_card;comment:开户许可证;"`
      MchName  string `json:"mchName" form:"mchName" gorm:"column:mch_name;comment:商户名称;"`
      MchForShort  string `json:"mchForShort" form:"mchForShort" gorm:"column:mch_for_short;comment:商户简称;"`
      MchAddress  string `json:"mchAddress" form:"mchAddress" gorm:"column:mch_address;comment:商户地址;"`
      LegalPerson  string `json:"legalPerson" form:"legalPerson" gorm:"column:legal_person;comment:与营业执照相同的法人名字;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:法人电话;"`
      LegalCard  string `json:"legalCard" form:"legalCard" gorm:"column:legal_card;comment:法人身份证号;"`
      CardFront  string `json:"cardFront" form:"cardFront" gorm:"column:card_front;comment:身份证正面;"`
      CardReverse  string `json:"cardReverse" form:"cardReverse" gorm:"column:card_reverse;comment:身份证反面;"`
      CardValidity  string `json:"cardValidity" form:"cardValidity" gorm:"column:card_validity;comment:身份证有效期;"`
      MchType  string `json:"mchType" form:"mchType" gorm:"column:mch_type;comment:商户类型;"`
      Status  int `json:"status" form:"status" gorm:"column:status;comment:商户认证状态;"`
      LinkmanName  string `json:"linkmanName" form:"linkmanName" gorm:"column:linkman_name;comment:联系人姓名;"`
      LinkmanPhone  string `json:"linkmanPhone" form:"linkmanPhone" gorm:"column:linkman_phone;comment:联系人电话;"`
      LinkmanMail  string `json:"linkmanMail" form:"linkmanMail" gorm:"column:linkman_mail;comment:联系人邮箱;"`
      LinkmanAddress  string `json:"linkmanAddress" form:"linkmanAddress" gorm:"column:linkman_address;comment:联系人通讯地址;"`
      BankFrontImg  string `json:"bankFrontImg" form:"bankFrontImg" gorm:"column:bank_front_img;comment:银行卡正面图片;"`
      CloseType  string `json:"closeType" form:"closeType" gorm:"column:close_type;comment:结算卡类型;"`
      CloseCardNum  string `json:"closeCardNum" form:"closeCardNum" gorm:"column:close_card_num;comment:结算帐户卡号;"`
      CloseCardName  string `json:"closeCardName" form:"closeCardName" gorm:"column:close_card_name;comment:结算帐户户名;"`
      BankPhone  string `json:"bankPhone" form:"bankPhone" gorm:"column:bank_phone;comment:银行预留手机号;"`
      OpenBank  string `json:"openBank" form:"openBank" gorm:"column:open_bank;comment:开户行;"`
      OpenBankNum  string `json:"openBankNum" form:"openBankNum" gorm:"column:open_bank_num;comment:开户行编号支行;"`
      OpenBankName  string `json:"openBankName" form:"openBankName" gorm:"column:open_bank_name;comment:开户行名称支行;"`
      OpenBankCity  string `json:"openBankCity" form:"openBankCity" gorm:"column:open_bank_city;comment:开户行省市区;"`
      StartMoney  float64 `json:"startMoney" form:"startMoney" gorm:"column:start_money;comment:起始结算金额;type:decimal(10,2);"`
      LegalPersonClose  string `json:"legalPersonClose" form:"legalPersonClose" gorm:"column:legal_person_close;comment:法人结算;"`
      NotLegalPersonName  string `json:"notLegalPersonName" form:"notLegalPersonName" gorm:"column:not_legal_person_name;comment:非法人姓名;"`
      NotLegalPersonCardNum  string `json:"notLegalPersonCardNum" form:"notLegalPersonCardNum" gorm:"column:not_legal_person_card_num;comment:非法人身份证号码;"`
      NotLegalPersonAuthImg  string `json:"notLegalPersonAuthImg" form:"notLegalPersonAuthImg" gorm:"column:not_legal_person_auth_img;comment:非法人结算授权书;"`
      AuthCardFrontImg  string `json:"authCardFrontImg" form:"authCardFrontImg" gorm:"column:auth_card_front_img;comment:被授权人身份证正面;"`
      AuthCardReverseImg  string `json:"authCardReverseImg" form:"authCardReverseImg" gorm:"column:auth_card_reverse_img;comment:被授权人身份证反面;"`
      Rate  string `json:"rate" form:"rate" gorm:"column:rate;comment:费率的所有信息;"`
      AgreementType  string `json:"agreementType" form:"agreementType" gorm:"column:agreement_type;comment:协议类型 电子 纸质;"`
      StoreHeadImg  string `json:"storeHeadImg" form:"storeHeadImg" gorm:"column:store_head_img;comment:门店门头照;"`
      StoreImg  string `json:"storeImg" form:"storeImg" gorm:"column:store_img;comment:门店内景照;"`
      StoreCashierImg  string `json:"storeCashierImg" form:"storeCashierImg" gorm:"column:store_cashier_img;comment:门店收银台照;"`
      ContractImg  string `json:"contractImg" form:"contractImg" gorm:"column:contract_img;comment:租赁合同;"`
      OtherImg  string `json:"otherImg" form:"otherImg" gorm:"column:other_img;comment:其它资料图片;"`
}


// TableName AppPay 表名
func (AppPay) TableName() string {
  return "app_pay"
}

