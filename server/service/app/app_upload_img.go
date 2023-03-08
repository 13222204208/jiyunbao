package app

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"
)

type AppUploadImgService struct {
}

// CreateAppUploadImg 创建AppUploadImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUploadImgService *AppUploadImgService) CreateAppUploadImg(appUploadImg app.AppUploadImg) (err error) {

	custInfoId := *appUploadImg.CustInfoId

	if err = IsSetImgType(custInfoId, *appUploadImg.PicType); err != nil {
		return err
	}

	appUploadImg.PicTypeName = transitionImgType(*appUploadImg.PicType)

	err = global.GVA_DB.Create(&appUploadImg).Error

	//if err == nil {
	//	err = SmscUpload(appUploadImg.ReqId, appUploadImg.SysFlowId, appUploadImg.PicTypeName, appUploadImg.ImgUrl)
	//}
	return err
}

//图片上送
func (appUploadImgService *AppUploadImgService) ImgSubmit(id uint) (err error) {
	var i app.AppUploadImg
	global.GVA_DB.Where("id = ?", id).First(&i)
	if i.ImgUrl == "" {
		return errors.New("查找不到数据")
	}
	custInfoId := *i.CustInfoId
	err, info, _ := GetCustInfo(uint(custInfoId))
	if err != nil {
		return errors.New("获取商户进件数据失败")
	}
	i.ReqId = info["reqId"]
	i.SysFlowId = info["sysFlowId"]
	global.GVA_DB.Model(&i).Where("id = ?", id).Updates(map[string]interface{}{"req_id": info["reqId"], "sys_flow_id": info["sysFlowId"]})
	err = SmscUpload(i.ReqId, i.SysFlowId, i.PicTypeName, i.ImgUrl)
	return err
}

// DeleteAppUploadImg 删除AppUploadImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUploadImgService *AppUploadImgService) DeleteAppUploadImg(appUploadImg app.AppUploadImg) (err error) {
	err = global.GVA_DB.Delete(&appUploadImg).Error
	return err
}

// DeleteAppUploadImgByIds 批量删除AppUploadImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUploadImgService *AppUploadImgService) DeleteAppUploadImgByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppUploadImg{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppUploadImg 更新AppUploadImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUploadImgService *AppUploadImgService) UpdateAppUploadImg(appUploadImg app.AppUploadImg) (err error) {

	appUploadImg.PicTypeName = transitionImgType(*appUploadImg.PicType)
	appUploadImg.ReqId = utils.ReqId()
	err = global.GVA_DB.Save(&appUploadImg).Error
	fmt.Println("保存结果", err)

	custInfoId := uint(*appUploadImg.CustInfoId)

	var c app.AppCustInfo
	err = global.GVA_DB.Where("id", custInfoId).First(&c).Error
	if err != nil && c.ChangeSysFlowId != "" {
		return errors.New("获取变更申请流水号失败")
	}

	err = ChangeSmscUpload(appUploadImg.ReqId, c.ChangeSysFlowId, appUploadImg.PicTypeName, appUploadImg.ImgUrl)

	return err
}

// GetAppUploadImg 根据id获取AppUploadImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUploadImgService *AppUploadImgService) GetAppUploadImg(id uint) (appUploadImg app.AppUploadImg, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appUploadImg).Error
	return
}

// GetAppUploadImgInfoList 分页获取AppUploadImg记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUploadImgService *AppUploadImgService) GetAppUploadImgInfoList(info appReq.AppUploadImgSearch) (list []app.AppUploadImg, total int64, err error) {
	fmt.Println("返回的数据ID", info.CustInfoId)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppUploadImg{})
	var appUploadImgs []app.AppUploadImg
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	db = db.Where("cust_info_id = ?", info.CustInfoId)
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appUploadImgs).Error
	return appUploadImgs, total, err
}

//判断此类型图片，是否已经提交
func IsSetImgType(custInfoId, picType int) error {
	if (!errors.Is(global.GVA_DB.Where("cust_info_id = ? AND pic_type = ?", custInfoId, picType).First(&app.AppUploadImg{}).Error, gorm.ErrRecordNotFound)) {
		return errors.New("已经存在相同类型的数据")
	} else {
		return nil
	}
}

//图片上送
func SmscUpload(reqId, sysFlowId, picTypeName, filePath string) error {
	err, url := GetSetUrl(3)
	if err != nil {
		return errors.New("获取图片上送接口失败")
	}

	info := map[string]string{
		"reqId":     reqId,
		"sysFlowId": sysFlowId,
		"picType":   picTypeName,
		"filePath":  filePath,
	}
	fmt.Println("提交的数据", info)
	err, res := YsPost(url, info)
	if err != nil {
		return err
	} else {
		err, result := TrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误")
		}
		fmt.Println("返回的结果", result)
		return CommonErrorCode(result)
	}
}

//变更图片上送
func ChangeSmscUpload(reqId, sysFlowId, picTypeName, filePath string) error {
	err, url := GetSetUrl(12)
	if err != nil {
		return errors.New("获取图片上送接口失败")
	}

	info := map[string]string{
		"reqId":        reqId,
		"changeFlowId": sysFlowId,
		"picType":      picTypeName,
		"filePath":     filePath,
	}
	err, res := YsPost(url, info)
	if err != nil {
		return err
	} else {
		err, result := TrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误")
		}
		fmt.Println("返回的状态", result.SubCode)
		return CommonErrorCode(result)
	}
}

//根据picType返回图片类别
//https://gateway-doc.ysepay.com/gatewayDocs/summary/N0000230/N0000231/N0000233/I0000175.html
//图片类别,A001-营业执照 A002-法人身份证正面(头像面) A003-法人身份证反面(国徽面) A004-结算账户正面(卡号面) A005-结算账户反面 A006-商户门头照片 A007-内景照片 A008-收银台照片 A009-手持身份证合影照片 A010-收单协议盖章页 A011-开户许可证 A012-收单协议首页 A013-非法人身份证头像面 A014-非法人身份证国徽面 B001-租赁合同 第一页 B002-租赁合同 第二页 B003-租赁合同 第三页 B004-法人/非法人手持授权书 B005-法人/非法人结算授权书 B006-租赁面积图片 B007-经营业务图片 B008-其他1 B009-其他2

func transitionImgType(picType int) string {
	switch {
	case picType == 1:
		return "A001"
	case picType == 2:
		return "A002"
	case picType == 3:
		return "A003"
	case picType == 4:
		return "A004"
	case picType == 5:
		return "A005"
	case picType == 6:
		return "A006"
	case picType == 7:
		return "A007"
	case picType == 8:
		return "A008"
	case picType == 9:
		return "A009"
	case picType == 10:
		return "A0010"
	case picType == 11:
		return "A0011"
	case picType == 12:
		return "A0012"
	case picType == 13:
		return "A0013"
	case picType == 14:
		return "A0014"
	case picType == 15:
		return "B001"
	case picType == 16:
		return "B002"
	case picType == 17:
		return "B003"
	case picType == 18:
		return "B004"
	case picType == 19:
		return "B005"
	case picType == 20:
		return "B006"
	case picType == 21:
		return "B007"
	case picType == 22:
		return "B008"
	case picType == 23:
		return "B009"
	default:
		return ""
	}
}
