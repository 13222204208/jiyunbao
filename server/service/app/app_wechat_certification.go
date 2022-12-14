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

type AppWechatCertificationService struct {
}

// CreateAppWechatCertification 创建AppWechatCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appWechatCertificationService *AppWechatCertificationService) CreateAppWechatCertification(appWechatCertification app.AppWechatCertification) (err error) {
	crpCertNo, cid := MercIdCustInfo(appWechatCertification.MercId)
	if crpCertNo != "" {
		appWechatCertification.ContactCorpId = crpCertNo
	}

	if cid != 0 {
		s := GetCustInfoImgUrl(cid, 2) //法人身份证正面
		if s != "" {
			appWechatCertification.SfzfrontPhoto = s
		}

		s = GetCustInfoImgUrl(cid, 3) //法人身份证反面
		if s != "" {
			appWechatCertification.SfzbackPhoto = s
		}

		s = GetCustInfoImgUrl(cid, 6) //门头照
		if s != "" {
			appWechatCertification.ShopPhoto = s
		}

		s = GetCustInfoImgUrl(cid, 7) //内景照片
		if s != "" {
			appWechatCertification.StoreEnvirPhoto = s
		}
	}
	err = global.GVA_DB.Create(&appWechatCertification).Error
	return err
}

// DeleteAppWechatCertification 删除AppWechatCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appWechatCertificationService *AppWechatCertificationService) DeleteAppWechatCertification(appWechatCertification app.AppWechatCertification) (err error) {
	err = global.GVA_DB.Delete(&appWechatCertification).Error
	return err
}

// DeleteAppWechatCertificationByIds 批量删除AppWechatCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appWechatCertificationService *AppWechatCertificationService) DeleteAppWechatCertificationByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppWechatCertification{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppWechatCertification 更新AppWechatCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appWechatCertificationService *AppWechatCertificationService) UpdateAppWechatCertification(appWechatCertification app.AppWechatCertification) (err error) {
	err = global.GVA_DB.Save(&appWechatCertification).Error
	return err
}

// GetAppWechatCertification 根据id获取AppWechatCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appWechatCertificationService *AppWechatCertificationService) GetAppWechatCertification(id uint) (appWechatCertification app.AppWechatCertification, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appWechatCertification).Error
	return
}

// GetAppWechatCertificationInfoList 分页获取AppWechatCertification记录
// Author [piexlmax](https://github.com/piexlmax)
func (appWechatCertificationService *AppWechatCertificationService) GetAppWechatCertificationInfoList(info appReq.AppWechatCertificationSearch) (list []app.AppWechatCertification, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppWechatCertification{})
	var appWechatCertifications []app.AppWechatCertification
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
	err = db.Limit(limit).Offset(offset).Find(&appWechatCertifications).Error
	return appWechatCertifications, total, err
}

//微信实名认证上送
func (appWechatCertificationService *AppWechatCertificationService) WechatCertification(id uint) (err error) {
	err, info := GetAppWechatCertificationInfo(id)
	if err != nil {
		return errors.New("获取微信认证数据失败")
	}
	//fmt.Println("提交的信息", info, "请求的url", url)
	err, url := GetSetUrl(9)
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
		if result.BusinessData.ApplyNo != "" {
			SaveApplyNo(id, result.BusinessData.ApplyNo)
		}
		return CommonErrorCode(result)
	}
}

//查询微信实名认证状态
func (appWechatCertificationService *AppWechatCertificationService) CertificationState(applyNo string) (err error) {

	err, url := GetSetUrl(10)
	if err != nil {
		return errors.New("获取上送接口失败")
	}
	info := map[string]string{
		"reqId":   utils.ReqId(),
		"applyNo": applyNo,
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
		return GetApplyState(result)
	}
}

//返回的微信实名认证申请状态
func GetApplyState(result YsResult) (err error) {
	code := result.Code
	subCode := result.SubCode
	fmt.Println("网关响应码", code, "业务响应码", subCode)
	switch {
	case subCode == "0000":
		applyState := result.BusinessData.ApplyState
		if applyState == "APPLYMENT_STATE_REJECTED" {
			return errors.New(result.BusinessData.RejectReason)
		}

		if applyState == "APPLYMENT_STATE_PASSED" {
			return nil
		} else {
			msg := GetApplyStateMsg(applyState)
			return errors.New(msg)
		}

	case code == "00000":
		return errors.New(result.SubMsg)
	default:
		return errors.New(result.Msg)
	}
}

//返回的微信实名认证申请状态消息转换
func GetApplyStateMsg(state string) string {
	switch {
	case state == "APPLYMENT_STATE_WAITTING_FOR_AUDIT":
		return "审核中，请耐心等待1~2个工作日，微信支付将会完成审核"
	case state == "APPLYMENT_STATE_EDITTING":
		return "编辑中，可能提交申请发生了错误导致，可用同一个业务申请编号重新提交"
	case state == "APPLYMENT_STATE_WAITTING_FOR_CONFIRM_CONTACT":
		return "待确认联系信息，请扫描微信支付返回的小程序码确认联系信息"
	case state == "APPLYMENT_STATE_WAITTING_FOR_CONFIRM_LEGALPERSON":
		return "待账户验证，请扫描微信支付返回的小程序码在小程序端完成账户验证"
	case state == "APPLYMENT_STATE_FREEZED":
		return "已冻结，可能是该主体已完成过入驻"
	case state == "APPLYMENT_STATE_CANCELED":
		return "已作废，表示申请单已被撤销，无需再对其进行操作"
	default:
		return "错误"
	}
}

//保存申请单编号
func SaveApplyNo(id uint, applyNo string) {
	var w app.AppWechatCertification
	err := global.GVA_DB.Model(&w).Where("id = ?", id).Update("apply_no", applyNo).Error
	if err != nil {
		global.GVA_LOG.Error("保存申请单编号失败!", zap.Error(err))
	}

}

//获取微信实名认证的信息
func GetAppWechatCertificationInfo(id uint) (err error, info map[string]string) {
	var w app.AppWechatCertification
	err = global.GVA_DB.Where("id = ?", id).First(&w).Error
	if err != nil {
		return
	}
	reqId := utils.ReqId()
	info = map[string]string{
		"reqId":           reqId,
		"mercId":          w.MercId,
		"mercName":        w.MercName,
		"merContactType":  strconv.Itoa(*w.MerContactType),
		"idTypeCd":        w.IdTypeCd,
		"idAddress":       w.IdAddress,
		"managementType":  "0" + strconv.Itoa(*w.ManagementType),
		"storeName":       w.StoreName,
		"idFrontImg":      w.IdFrontImg,
		"idBackImg":       w.IdBackImg,
		"contactCorpId":   w.ContactCorpId,
		"shopPhoto":       w.ShopPhoto,
		"storeEnvirPhoto": w.StoreEnvirPhoto,
	}
	global.GVA_DB.Model(&w).Where("id = ?", id).Update("req_id", reqId)
	return err, info
}

//根据商户号获取商户的认证信息
func MercIdCustInfo(mercId string) (str string, cid uint) {
	var c app.AppSignContract
	err := global.GVA_DB.Where("merc_id", mercId).First(&c).Error
	if err != nil {
		return "", cid
	} else {
		var custInfo app.AppCustInfo
		err = global.GVA_DB.Where("cust_id = ?", c.CustId).First(&custInfo).Error
		if err != nil {
			return "", cid
		} else {
			return custInfo.CrpCertNo, custInfo.ID
		}
	}
}

//根据custId 获取上传的图片地址，然后转为 base64格式
func GetCustInfoImgUrl(cid uint, picType int) string {
	var img app.AppUploadImg
	err := global.GVA_DB.Where("cust_info_id = ? AND pic_type = ?", cid, picType).First(&img).Error
	fmt.Println("获取的图片数据", img)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		return ""
	} else {
		_, baseImg := UploadDocument(img.ImgUrl)
		if baseImg != "" {
			return baseImg
		} else {
			return ""
		}
	}
}

//文件上传
func UploadDocument(imgUrl string) (err error, image string) {
	err, url := GetSetUrl(11)
	if err != nil {
		return errors.New("获取上送接口失败"), ""
	}
	info := map[string]string{
		"reqId":    utils.ReqId(),
		"filePath": imgUrl,
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
			return err, result.BusinessData.Url
		} else {
			return err, ""
		}
	}
}

//微信实名认证明细查询
func (appWechatCertificationService *AppWechatCertificationService) GetAuthMessagesByMercId(mercId string) (err error, status string) {
	err, url := GetSetUrl(14)
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
		//fmt.Println("返回的数据", result.BusinessData)
		if result.SubCode == "0000" {
			var applyStatus string
			if result.BusinessData.OutMercIdRemark != "" {
				applyStatus = result.BusinessData.OutMercIdRemark
			} else {
				applyStatus = ApplyStatusMsg(result.BusinessData.ApplyStatus)
			}
			if result.BusinessData.QrcodeData != "" {
				UpdatesApplyInfo(mercId, map[string]interface{}{
					"qrcode_data":                result.BusinessData.QrcodeData,
					"send_channel_auth_apply_id": result.BusinessData.SendChannelAuthApplyId,
				})
			}
			return err, applyStatus
		} else {
			return err, ""
		}
	}
}

//实名认证授权状态
func (appWechatCertificationService *AppWechatCertificationService) GetAuthState(mercId string) (err error, status string) {
	err, url := GetSetUrl(15)
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
			thirdMercIdX := result.BusinessData.ThirdMercIdX
			if thirdMercIdX != "" {
				UpdateThirdMercIdX(mercId, thirdMercIdX)
			}
			return err, AuthStateXMsg(result.BusinessData.AuthStateX)
		} else {
			return err, ""
		}
	}
}

//保存第三方商户号
func UpdateThirdMercIdX(mercId, thirdMercIdX string) {
	var c app.AppWechatCertification
	err := global.GVA_DB.Model(&c).Where("merc_id", mercId).Update("third_merc_id_x", thirdMercIdX).Error
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
	}
}

//保存小程序码图片,及信息
func UpdatesApplyInfo(mercId string, info map[string]interface{}) {
	var c app.AppWechatCertification
	err := global.GVA_DB.Model(&c).Where("merc_id", mercId).Updates(info).Error
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
	}
}

//授权状态转换
func AuthStateXMsg(state string) string {
	switch state {
	case "AUTHORIZE_STATE_UNAUTHORIZED":
		return "未授权"
	case "AUTHORIZE_STATE_AUTHORIZED":
		return "已授权"
	default:
		return "转换失败"
	}
}

//申请状态转换
func ApplyStatusMsg(status string) string {
	switch status {
	case "00":
		return "审核中"
	case "01":
		return "待确认联系信息"
	case "02":
		return "待账户验证"
	case "03":
		return "审核通过"
	case "04":
		return "审核驳回"
	case "05":
		return "冻结"
	case "06":
		return "已作废"
	case "07":
		return "申请单提交失败"
	default:
		return "返回失败"
	}
}

//根据custId 获取上传的图片地址，然后转为 base64格式
//func GetCustInfoImgBase64(cid uint, picType int) string {
//	var img app.AppUploadImg
//	err := global.GVA_DB.Where("cust_info_id = ? AND pic_type = ?", cid, picType).First(&img).Error
//	fmt.Println("获取的图片数据", img)
//	if err != nil {
//		global.GVA_LOG.Error("获取失败!", zap.Error(err))
//		return ""
//	} else {
//		baseImg, err := GetUrlImgBase64(img.ImgUrl)
//		if err != nil {
//			return ""
//		} else {
//			return baseImg
//		}
//	}
//}
//
//func GetUrlImgBase64(path string) (baseImg string, err error) {
//	str, _ := os.Getwd()
//	path = str + "/" + path
//	fmt.Println("当前路径", path)
//	//获取本地文件
//	file, err := os.Open(path)
//	if err != nil {
//		err = errors.New("获取本地图片失败")
//		return
//	}
//	defer file.Close()
//	imgByte, _ := ioutil.ReadAll(file)
//
//	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
//	//取图片类型
//	mimeType := http.DetectContentType(imgByte)
//	switch mimeType {
//	case "image/jpeg":
//		baseImg = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte)
//	case "image/png":
//		baseImg = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
//	}
//	return
//}
