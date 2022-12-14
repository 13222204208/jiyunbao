// 自动生成模板AppCustInfo
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppCustInfo 结构体
type AppCustInfo struct {
	global.GVA_MODEL
	Uid             *int   `json:"uid" form:"uid" gorm:"column:uid;comment:提交 的商户ID;"`
	ReqId           string `json:"reqId" form:"reqId" gorm:"column:req_id;comment:要求32个字符内（最少14个字符）;size:32;"`
	CertId          string `json:"certId" form:"certId" gorm:"column:cert_id;comment:发起方商户号，服务商在银盛给自己开设的商户号;"`
	AgtMercId       string `json:"agtMercId" form:"agtMercId" gorm:"column:agt_merc_id;comment:代理商编号,默认同发起方代理商编号;"`
	MercName        string `json:"mercName" form:"mercName" gorm:"column:merc_name;comment:商户名称;"`
	MercShortName   string `json:"mercShortName" form:"mercShortName" gorm:"column:merc_short_name;comment:商户简称,默认为商户名称，若商户名称长度超过20,商户简称必传;"`
	MercType        *int   `json:"mercType" form:"mercType" gorm:"column:merc_type;comment:商户类型,2 小微 3个体 4企业;"`
	MccCd           string `json:"mccCd" form:"mccCd" gorm:"column:mcc_cd;comment:通过"查询MCC码"接口获取mccCd;"`
	ContactMail     string `json:"contactMail" form:"contactMail" gorm:"column:contact_mail;comment:;"`
	ContactMan      string `json:"contactMan" form:"contactMan" gorm:"column:contact_man;comment:联系人,默认同法人;"`
	ContactPhone    string `json:"contactPhone" form:"contactPhone" gorm:"column:contact_phone;comment:联系人电话, 默认同法人电话号码;"`
	CusMgrNm        string `json:"cusMgrNm" form:"cusMgrNm" gorm:"column:cus_mgr_nm;comment:客户经理;"`
	NotifyUrl       string `json:"notifyUrl" form:"notifyUrl" gorm:"column:notify_url;comment:异步通知地址,银盛通过商户提供的地址下发通知;"`
	CrpCertNo       string `json:"crpCertNo" form:"crpCertNo" gorm:"column:crp_cert_no;comment:法人证件号;size:30;"`
	OpenCertNo      string `json:"openCertNo" form:"openCertNo" gorm:"column:open_cert_no;comment:开开户证件号,默认法人证件号,非法人结算必填"`
	CrpCertType     string `json:"crpCertType" form:"crpCertType" gorm:"column:crp_cert_type;default:00;comment:法人证件类型,暂时只支持 00 身份证;"`
	CertBgn         string `json:"certBgn" form:"certBgn" gorm:"column:cert_bgn;comment:证件开始日期,证件开始日期yyyyMMdd;size:8;"`
	CertExpire      string `json:"certExpire" form:"certExpire" gorm:"column:cert_expire;comment:证件有效期,日期格式yyyyMMdd ，如果为长期或者永久，请填值“29991231”;"`
	CrpNm           string `json:"crpNm" form:"crpNm" gorm:"column:crp_nm;comment:法人姓名;"`
	CrpPhone        string `json:"crpPhone" form:"crpPhone" gorm:"column:crp_phone;comment:法人手机号;"`
	StlAccNo        string `json:"stlAccNo" form:"stlAccNo" gorm:"column:stl_acc_no;comment:结算账号;size:30;"`
	BankSubCode     string `json:"bankSubCode" form:"bankSubCode" gorm:"column:bank_sub_code;comment:开户支行联行号,通过"行名行号查询"获取bankCod;size:20;"`
	StlAccType      *int   `json:"stlAccType" form:"stlAccType" gorm:"column:stl_acc_type;comment:,11 对私 21 对公 23 对公存折 24 单位结算卡;"`
	BusProviceCode  string `json:"busProviceCode" form:"busProviceCode" gorm:"column:bus_provice_code;comment:营业归属省代码,通过"地区信息查询"接口获取cityCd;"`
	BusCityCode     string `json:"busCityCode" form:"busCityCode" gorm:"column:bus_city_code;comment:营业归属市代码,通过"地区信息查询"接口获取cityCd;"`
	BusAreaCode     string `json:"busAreaCode" form:"busAreaCode" gorm:"column:bus_area_code;comment:营业归属区(县)代码,通过接口获取cityCd，详见地区信息查询，城市下有区(县必填)，城市下无区(县)则不上送，如广东省中山市，广东省东莞市"`
	BusAddr         string `json:"busAddr" form:"busAddr" gorm:"column:bus_addr;comment:包含省市区，示例：广东省深圳市龙华新区大道513号远景大厦;"`
	SysFlowId       string `json:"sysFlowId" form:"sysFlowId" gorm:"column:sys_flow_id;comment:入网申请流水资料上送成功后返回"`
	ChangeSysFlowId string `json:"changeSysFlowId" form:"changeSysFlowId" gorm:"column:change_sys_flow_id;comment:变更申请流水号"`
	CustId          string `json:"custId" form:"custId" gorm:"column:cust_id;comment:客户号"`
}

// TableName AppCustInfo 表名
func (AppCustInfo) TableName() string {
	return "app_cust_info"
}
