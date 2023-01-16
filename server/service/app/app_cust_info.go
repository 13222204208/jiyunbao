package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"regexp"
	"strconv"
)

type AppCustInfoService struct {
}

var ReqId = utils.ReqId()

// CreateAppCustInfo 创建AppCustInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCustInfoService *AppCustInfoService) CreateAppCustInfo(appCustInfo app.AppCustInfo) (err error) {
	appCustInfo.CrpCertType = "00"
	appCustInfo.NotifyUrl = GetAppSet().NotifyUrl
	appCustInfo.ReqId = ReqId //流水号
	err = global.GVA_DB.Create(&appCustInfo).Error
	return err
}

//app 端提交的支付认证
func (appCustInfoService *AppCustInfoService) Attestation(appCustInfo app.AppCustInfo) (err error) {
	err, _ = VerifyAppCustAttestation(appCustInfo.Uid)
	if err != nil {
		return err
	}
	appCustInfo.NotifyUrl = GetAppSet().NotifyUrl
	appCustInfo.ReqId = ReqId //流水号
	err = global.GVA_DB.Create(&appCustInfo).Error
	if err == nil {
		SaveToUploadImage(&appCustInfo)
	}
	return err
}

// DeleteAppCustInfo 删除AppCustInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCustInfoService *AppCustInfoService) DeleteAppCustInfo(appCustInfo app.AppCustInfo) (err error) {
	err = global.GVA_DB.Delete(&appCustInfo).Error
	return err
}

// DeleteAppCustInfoByIds 批量删除AppCustInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCustInfoService *AppCustInfoService) DeleteAppCustInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppCustInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppCustInfo 更新AppCustInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCustInfoService *AppCustInfoService) UpdateAppCustInfo(appCustInfo app.AppCustInfo) (err error) {
	appCustInfo.ReqId = ReqId
	err = global.GVA_DB.Save(&appCustInfo).Error
	return err
}

// GetAppCustInfo 根据id获取AppCustInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCustInfoService *AppCustInfoService) GetAppCustInfo(id uint) (appCustInfo app.AppCustInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appCustInfo).Error
	return
}

// GetAppCustInfoInfoList 分页获取AppCustInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCustInfoService *AppCustInfoService) GetAppCustInfoInfoList(info appReq.AppCustInfoSearch) (list []app.AppCustInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppCustInfo{})
	var appCustInfos []app.AppCustInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ReqId != "" {
		db = db.Where("req_id = ?", info.ReqId)
	}
	if info.MercName != "" {
		db = db.Where("merc_name LIKE ?", "%"+info.MercName+"%")
	}
	if info.MercType != nil {
		db = db.Where("merc_type = ?", info.MercType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appCustInfos).Error
	return appCustInfos, total, err
}

//获取商户进件的数据
func GetCustInfo(id uint) (err error, info map[string]string) {
	var appCustInfo app.AppCustInfo
	fmt.Println("资料ID", id)
	err = global.GVA_DB.Where("id = ?", id).First(&appCustInfo).Error
	if err != nil {
		return
	}

	ReqId = utils.ReqId() //流水号

	info = map[string]string{
		"reqId":          ReqId,
		"mercName":       appCustInfo.MercName,
		"mercShortName":  appCustInfo.MercShortName,
		"mercType":       strconv.Itoa(*appCustInfo.MercType),
		"mccCd":          appCustInfo.MccCd,
		"contactMail":    appCustInfo.ContactMail,
		"cusMgrNm":       appCustInfo.CusMgrNm,
		"notifyUrl":      appCustInfo.NotifyUrl,
		"crpCertNo":      appCustInfo.CrpCertNo,
		"crpCertType":    appCustInfo.CrpCertType,
		"certBgn":        appCustInfo.CertBgn,
		"certExpire":     appCustInfo.CertExpire,
		"crpNm":          appCustInfo.CrpNm,
		"crpPhone":       appCustInfo.CrpPhone,
		"stlAccNo":       appCustInfo.StlAccNo,
		"bankSubCode":    appCustInfo.BankSubCode,
		"openCertNo":     appCustInfo.OpenCertNo,
		"stlAccType":     strconv.Itoa(*appCustInfo.StlAccType),
		"busNo":          appCustInfo.BusNo,
		"busNm":          appCustInfo.BusNm,
		"busCertBgn":     appCustInfo.BusCertBgn,
		"busCertExpire":  appCustInfo.BusCertExpire,
		"busProviceCode": appCustInfo.BusProviceCode,
		"busCityCode":    appCustInfo.BusCityCode,
		"busAreaCode":    appCustInfo.BusAreaCode,
		"busAddr":        appCustInfo.BusAddr,
	}
	if appCustInfo.SysFlowId != "" {
		info["sysFlowId"] = appCustInfo.SysFlowId
	}
	fmt.Println("资料信息sysFlowId:", appCustInfo.SysFlowId)

	return
}

//去掉返回结果json数据中的 pre标签
func TrimPrefix(data string) (err error, s YsResult) {

	fmt.Println("最新返回数据", data)
	re, _ := regexp.Compile("\\<pre[\\S\\s]+?\\</pre\\>")
	src := re.ReplaceAllString(data, "")
	fmt.Println("去掉pre", src)
	err = json.Unmarshal([]byte(src), &s)
	if err != nil {
		fmt.Println("解析错误的信息: %v", err)
	}
	return err, s
}

//银盛公共返回参数
//code 网关响应码，示例值：00000 业务响应码，参见具体的API接口文档
//isCode 是网关响应码，还是业务响应码
//res
func CommonErrorCode(result YsResult) error {
	code := result.Code
	subCode := result.SubCode
	fmt.Println("网关响应码", code, "业务响应码", subCode)
	switch {
	case subCode == "0000":
		return nil
	case code == "00000":
		return errors.New(result.SubMsg)
	default:
		return errors.New(result.Msg)
	}
}

//资料上送  id 上送资料的ID
func (appCustInfoService *AppCustInfoService) AddCustInfoApply(id uint) (err error) {

	err, info := GetCustInfo(id)
	if err != nil {
		return errors.New("获取商户进件数据失败")
	}
	//fmt.Println("提交的信息", info, "请求的url", url)
	err, url := GetSetUrl(1)
	if err != nil {
		return errors.New("获取上送接口失败")
	}
	fmt.Println("上传资料信息", info)
	err, res := YsPost(url, info)
	if err != nil {
		return err
	} else {
		err, result := TrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误")
		}
		fmt.Println("返回的结果", result)
		if result.BusinessData.SysFlowId != "" {
			SaveSysFlowId(id, result.BusinessData.SysFlowId)
		}
		return CommonErrorCode(result)
	}
}

//保存上传资料的 sysFlowId
func SaveSysFlowId(custInfoId uint, sysFlowId string) {
	err := global.GVA_DB.Model(&app.AppCustInfo{}).Where("id = ?", custInfoId).Update("sys_flow_id", sysFlowId).Error
	if err != nil {
		fmt.Println("保存sysFlowId失败", err.Error())
	} else {
		fmt.Println("保存成功")
	}
}

//保存客户号 CustId
func SaveCustId(custInfoId uint, custId string) {
	err := global.GVA_DB.Model(&app.AppCustInfo{}).Where("id = ?", custInfoId).Update("cust_id", custId).Error
	if err != nil {
		fmt.Println("保存custId失败", err.Error())
	} else {
		fmt.Println("保存custId成功")
	}
}

//资料修改
func (appCustInfoService *AppCustInfoService) ModifyCustInfoApply(id uint) (err error) {
	err, info := GetCustInfo(id)
	if err != nil {
		return errors.New("获取商户进件数据失败")
	}
	//fmt.Println("提交的信息", info, "请求的url", url)
	err, url := GetSetUrl(2)
	if err != nil {
		return errors.New("获取上送接口失败")
	}

	err, res := YsPost(url, info)
	if err != nil {
		return err
	} else {
		err, result := TrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误")
		}
		fmt.Println("返回的结果", result, result.BusinessData.SysFlowId)
		if result.BusinessData.SysFlowId != "" {
			SaveSysFlowId(id, result.BusinessData.SysFlowId)
		}

		return CommonErrorCode(result)
	}
}

//资料确认
func (appCustInfoService *AppCustInfoService) AuditCustInfoApply(id uint) (err error) {
	err, info := GetCustInfo(id)
	if err != nil {
		return errors.New("获取商户进件数据失败")
	}
	err, url := GetSetUrl(4)
	if err != nil {
		return errors.New("获取上送接口失败")
	}

	info["reqId"] = utils.ReqId()
	info["auditFlag"] = "Y"
	err, res := YsPost(url, info)
	if err != nil {
		return err
	} else {
		err, result := TrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误")
		}
		fmt.Println("返回的结果", result, result.BusinessData.SysFlowId)
		custId := result.BusinessData.CustId
		if custId != "" {
			SaveCustId(id, custId)
		}
		return CommonErrorCode(result)
	}
}

//商户状态查询
func (appCustInfoService *AppCustInfoService) QueryCustApply(id uint) (err error) {
	err, info := GetCustInfo(id)
	if err != nil {
		return errors.New("获取商户进件数据失败")
	}
	err, url := GetSetUrl(5)
	if err != nil {
		return errors.New("获取上送接口失败")
	}

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

//基本信息变更申请
func (appCustInfoService *AppCustInfoService) ChangeMercBaseInfo(custId string) (err error) {

	err, url := GetSetUrl(13)
	if err != nil {
		return errors.New("获取上送接口失败")
	}
	info := map[string]string{
		"reqId":  utils.ReqId(),
		"custId": custId,
	}
	err, res := YsPost(url, info)
	if err != nil {
		return err
	} else {
		err, result := TrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误")
		}
		fmt.Println("返回的结果", result)
		if result.BusinessData.ChangeSysFlowId != "" {
			saveChangeSysFlowId(custId, result.BusinessData.ChangeSysFlowId)
		}
		return CommonErrorCode(result)
	}
}

//保存变更申请流水号
func saveChangeSysFlowId(custId, changeSysFlowId string) {
	err := global.GVA_DB.Model(&app.AppCustInfo{}).Where("cust_id = ?", custId).Update("change_sys_flow_id", changeSysFlowId).Error
	if err != nil {
		global.GVA_LOG.Error("保存变更申请流水号失败!", zap.Error(err))
	}
}

//验证用户是否已经提交支付认证
func VerifyAppCustAttestation(uid uint) (err error, c app.AppCustInfo) {
	if !errors.Is(global.GVA_DB.Where("uid = ?", uid).First(&c).Error, gorm.ErrRecordNotFound) {
		return errors.New("已经提交认证"), c
	}
	return
}

//保存图片到upload_img
func SaveToUploadImage(c *app.AppCustInfo) {
	fmt.Println("图片", c.BusImage)
	cid := int(c.ID)
	if c.BusImage != "" {
		SaveImage(c.BusImage, cid, 1, "A001")
	}

	if c.CardFront != "" {
		SaveImage(c.CardFront, cid, 2, "A002")
	}

	if c.CardReverse != "" {
		SaveImage(c.CardReverse, cid, 3, "A003")
	}

	if c.BankFrontImg != "" {
		SaveImage(c.BankFrontImg, cid, 4, "A004")
	}

	if c.BankReverseImg != "" {
		SaveImage(c.BankReverseImg, cid, 5, "A005")
	}

	if c.StoreHeadImg != "" {
		SaveImage(c.StoreHeadImg, cid, 6, "A006")
	}

	if c.StoreImg != "" {
		SaveImage(c.StoreImg, cid, 7, "A007")
	}

	if c.StoreCashierImg != "" {
		SaveImage(c.StoreCashierImg, cid, 8, "A008")
	}

}

func SaveImage(img string, cid, t int, n string) {
	var i app.AppUploadImg
	i.ImgUrl = img
	i.PicType = &t
	i.PicTypeName = n
	i.CustInfoId = &cid
	global.GVA_DB.Create(&i)
}
