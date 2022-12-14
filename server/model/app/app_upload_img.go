// 自动生成模板AppUploadImg
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppUploadImg 结构体
type AppUploadImg struct {
	global.GVA_MODEL
	CustInfoId  *int   `json:"custInfoId" form:"custInfoId" gorm:"column:cust_info_id;comment:商户进件编号ID;"`
	ReqId       string `json:"reqId" form:"reqId" gorm:"column:req_id;comment:流水号;"`
	PicType     *int   `json:"picType" form:"picType" gorm:"column:pic_type;comment:图片类别;"`
	PicTypeName string `json:"picTypeName" form:"picTypeName" gorm:"column:pic_type_name;comment:图片类别名称"`
	ImgUrl      string `json:"imgUrl" form:"imgUrl" gorm:"column:img_url;comment:;size:300;"`
	SysFlowId   string `json:"sysFlowId" form:"sysFlowId" gorm:"column:sys_flow_id;comment:入网申请流水号;"`
}

// TableName AppUploadImg 表名
func (AppUploadImg) TableName() string {
	return "app_upload_img"
}
