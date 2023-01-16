package app

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"strconv"
)

type AppAlipayCertificationService struct {
}

// CreateAppAlipayCertification 创建AppAlipayCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAlipayCertificationService *AppAlipayCertificationService) CreateAppAlipayCertification(appAlipayCertification app.AppAlipayCertification) (err error) {
	err, custInfo := AlipayMercIdCustInfo(appAlipayCertification.MercId)
	if custInfo.CrpCertNo != "" {
		appAlipayCertification.ContactCorpId = custInfo.CrpCertNo
		appAlipayCertification.ContactsTel = custInfo.CrpPhone
	}

	err = global.GVA_DB.Create(&appAlipayCertification).Error
	return err
}

// DeleteAppAlipayCertification 删除AppAlipayCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAlipayCertificationService *AppAlipayCertificationService) DeleteAppAlipayCertification(appAlipayCertification app.AppAlipayCertification) (err error) {
	err = global.GVA_DB.Delete(&appAlipayCertification).Error
	return err
}

// DeleteAppAlipayCertificationByIds 批量删除AppAlipayCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAlipayCertificationService *AppAlipayCertificationService) DeleteAppAlipayCertificationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppAlipayCertification{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppAlipayCertification 更新AppAlipayCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAlipayCertificationService *AppAlipayCertificationService) UpdateAppAlipayCertification(appAlipayCertification app.AppAlipayCertification) (err error) {
	err = global.GVA_DB.Save(&appAlipayCertification).Error
	return err
}

// GetAppAlipayCertification 根据id获取AppAlipayCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAlipayCertificationService *AppAlipayCertificationService) GetAppAlipayCertification(id uint) (appAlipayCertification app.AppAlipayCertification, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appAlipayCertification).Error
	return
}

// GetAppAlipayCertificationInfoList 分页获取AppAlipayCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAlipayCertificationService *AppAlipayCertificationService) GetAppAlipayCertificationInfoList(info appReq.AppAlipayCertificationSearch) (list []app.AppAlipayCertification, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppAlipayCertification{})
	var appAlipayCertifications []app.AppAlipayCertification
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.MercId != "" {
		db = db.Where("merc_id = ?", info.MercId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appAlipayCertifications).Error
	return appAlipayCertifications, total, err
}

//支付宝实名认证上送
func (appAlipayCertificationService *AppAlipayCertificationService) AlipayCertification(id uint) (err error) {
	err, info := GetAppAlipayCertificationInfo(id)
	if err != nil {
		return errors.New("获取支付宝实名认证数据失败")
	}
	//fmt.Println("提交的信息", info, "请求的url", url)
	err, url := GetSetUrl(18)
	if err != nil {
		return errors.New("获取上送接口失败")
	}
	err, res := YsPost(url, info)
	if err != nil {
		return err
	} else {
		fmt.Println("解析前的数据", res)
		err, result := TrimPrefix(res)
		fmt.Println("解析后的结果", result.BusinessData)
		if err != nil {
			return errors.New("解析结果错误")
		}
		if result.BusinessData.SendChannelApplyNo != "" {
			SaveSendChannelApplyNo(id, result.BusinessData.SendChannelApplyNo)
		}
		return CommonErrorCode(result)
	}
}

//查询支付宝实名认证状态
func (appAlipayCertificationService *AppAlipayCertificationService) AlipayCertificationState(mercId, applyNo string) (err error) {

	err, url := GetSetUrl(19)
	if err != nil {
		return errors.New("获取上送接口失败")
	}
	info := map[string]string{
		"reqId":              utils.ReqId(),
		"mercId":             mercId,
		"sendChannelApplyNo": applyNo,
	}
	err, res := YsPost(url, info)
	if err != nil {
		return err
	} else {
		err, result := TrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误")
		}
		fmt.Println("返回的数据", result.BusinessData)
		return GetAlipayCertificationState(result, mercId)
	}
}

//查询支付宝实名认证授权状态
//实名认证授权状态
func (appAlipayCertificationService *AppAlipayCertificationService) GetAlipayAuthState(mercId string) (err error, status string) {
	err, url := GetSetUrl(20)
	if err != nil {
		return errors.New("获取上送接口失败"), ""
	}
	info := map[string]string{
		"reqId":  utils.ReqId(),
		"mercId": mercId,
	}
	err, res := YsPost(url, info)
	if err != nil {
		return err, ""
	} else {
		err, result := TrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误"), ""
		}
		fmt.Println("返回的数据", result.BusinessData)
		if result.SubCode == "0000" {
			authState := result.BusinessData.AuthState
			if authState == "0" {
				status = "未授权"
			} else if authState == "1" {
				status = "已授权"
			} else if authState == "2" {
				status = "已销户"
			}
			return err, status
		} else {
			return err, ""
		}
	}
}

func GetAlipayCertificationState(result YsResult, mercId string) (err error) {
	code := result.Code
	subCode := result.SubCode
	fmt.Println("网关响应码", code, "业务响应码", subCode)
	switch {
	case subCode == "0000":
		applyState := result.BusinessData.ApplyState
		if applyState == "AUDIT_REJECT" {
			return errors.New(result.BusinessData.RejectReason)
		}

		if applyState == "CONTACT_CONFIRM" {
			SaveAlipayQrcodeData(mercId, result.BusinessData.QrcodeData)
			return errors.New("待确认联系信息")
		}

		if applyState == "AUDIT_PASS" {
			return nil
		} else {
			return errors.New(result.BusinessData.ApplyState)
		}

	case code == "00000":
		return errors.New(result.SubMsg)
	default:
		return errors.New(result.Msg)
	}
}

//保存支付宝实名认证二维码
func SaveAlipayQrcodeData(mercId, qrcode string) {
	var a app.AppAlipayCertification
	err := global.GVA_DB.Model(&a).Where("merc_id = ?", mercId).Update("qrcode_data", qrcode).Error
	if err != nil {
		fmt.Println("错误信息", err.Error())
	}
}

//获取支付宝实名认证信息
func GetAppAlipayCertificationInfo(id uint) (err error, info map[string]string) {
	var w app.AppAlipayCertification
	err = global.GVA_DB.Where("id = ?", id).First(&w).Error
	if err != nil {
		return
	}
	reqId := utils.ReqId()
	info = map[string]string{
		"reqId":          reqId,
		"mercId":         w.MercId,
		"idTypeCd":       "0" + w.IdTypeCd,
		"storeName":      w.StoreName,
		"contactCorpId":  w.ContactCorpId,
		"contactsTel":    w.ContactsTel,
		"merContactType": strconv.Itoa(*w.MerContactType),
	}
	global.GVA_DB.Model(&w).Where("id = ?", id).Update("req_id", reqId)
	return err, info
}

//根据商户号获取商户的认证信息
func AlipayMercIdCustInfo(mercId string) (err error, custInfo app.AppCustInfo) {
	var c app.AppSignContract
	err = global.GVA_DB.Where("merc_id", mercId).First(&c).Error
	if err != nil {
		return err, custInfo
	} else {
		err = global.GVA_DB.Where("cust_id = ?", c.CustId).First(&custInfo).Error
		return err, custInfo

	}
}

//保存支付宝申请单编号
func SaveSendChannelApplyNo(id uint, applyNo string) {
	var w app.AppAlipayCertification
	err := global.GVA_DB.Model(&w).Where("id = ?", id).Update("send_channel_apply_no", applyNo).Error
	if err != nil {
		global.GVA_LOG.Error("保存申请单编号失败!", zap.Error(err))
	}
}
