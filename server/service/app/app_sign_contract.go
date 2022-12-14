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
	"strconv"
)

type AppSignContractService struct {
}

// CreateAppSignContract 创建AppSignContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSignContractService *AppSignContractService) CreateAppSignContract(appSignContract app.AppSignContract) (err error) {
	err = global.GVA_DB.Create(&appSignContract).Error
	return err
}

// DeleteAppSignContract 删除AppSignContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSignContractService *AppSignContractService) DeleteAppSignContract(appSignContract app.AppSignContract) (err error) {
	err = global.GVA_DB.Delete(&appSignContract).Error
	return err
}

// DeleteAppSignContractByIds 批量删除AppSignContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSignContractService *AppSignContractService) DeleteAppSignContractByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppSignContract{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppSignContract 更新AppSignContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSignContractService *AppSignContractService) UpdateAppSignContract(appSignContract app.AppSignContract) (err error) {
	//err = global.GVA_DB.Save(&appSignContract).Error
	//
	//return err
	appSignContract.ReqId = utils.ReqId()
	appSignContract.NotifyUrl = GetAppSet().NotifyUrl
	global.GVA_DB.Save(&appSignContract)
	err = SignContract(appSignContract)
	return err
}

// GetAppSignContract 根据id获取AppSignContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSignContractService *AppSignContractService) GetAppSignContract(id uint) (appSignContract app.AppSignContract, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appSignContract).Error
	return
}

// GetAppSignContractInfoList 分页获取AppSignContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSignContractService *AppSignContractService) GetAppSignContractInfoList(info appReq.AppSignContractSearch) (list []app.AppSignContract, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppSignContract{})
	var appSignContracts []app.AppSignContract
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appSignContracts).Error
	return appSignContracts, total, err
}

//合同签约申请
func SignContract(appSignContract app.AppSignContract) (err error) {
	info := map[string]string{
		"reqId":        appSignContract.ReqId,
		"busOpenType":  appSignContract.BusOpenType,
		"contractType": strconv.Itoa(*appSignContract.ContractType),
		"isSendConMsg": strconv.Itoa(*appSignContract.IsSendConMsg),
		"custId":       appSignContract.CustId,
		"notifyUrl":    appSignContract.NotifyUrl,
	}

	err, url := GetSetUrl(7)
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
		if result.SubCode == "0000" {
			UpdateSignContract(appSignContract.ID, result.BusinessData)
		}

		return CommonErrorCode(result)
	}
}

//修改合同签约申请内容
func UpdateSignContract(id uint, b BusinessData) {
	var c app.AppSignContract
	err := global.GVA_DB.Model(&c).Where("id = ?", id).Updates(map[string]interface{}{
		"sign_url": b.SignUrl,
		"sign_id":  b.SignId,
		"auth_id":  b.AuthId,
	}).Error
	if err != nil {
		fmt.Println("修改合同签约失败", err.Error())
	} else {
		fmt.Println("修改合同签约成功")
	}
}

//签约状态查询
func (appSignContractService *AppSignContractService) QueryAuthInfo(s app.AppSignContract) (error, string) {
	err, url := GetSetUrl(8)
	if err != nil {
		return errors.New("获取上送接口失败"), ""
	}
	info := map[string]string{
		"reqId":  utils.ReqId(),
		"authId": s.AuthId,
	}
	fmt.Println("签约状态数据", info)
	err, res := YsPost(url, info)
	if err != nil {
		return err, ""
	} else {
		err, result := TrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误"), ""
		}
		fmt.Println("返回的结果", result)

		str := result.BusinessData.Status
		switch str {
		case "00":
			str = "成功"
			UpdateMercId(s.ID, result.BusinessData.MercId)
		case "01":
			str = "初始化"
		case "02":
			str = "签约中"
		case "03":
			str = "待审核"
		case "04":
			str = "审核拒绝"
		}

		return CommonErrorCode(result), str
	}
}

//商户余额查询
func (appSignContractService *AppSignContractService) BalanceQuery(mercId string) (err error, total string) {
	err, url := GetSetUrl(17)
	if err != nil {
		return errors.New("获取上送接口失败"), ""
	}
	info := map[string]string{
		"merchant_usercode": mercId,
	}
	err, res := YsPost(url, info)
	fmt.Println("返回的结果", res)
	if err != nil {
		return err, ""
	} else {
		err, result := BalanceQueryTrimPrefix(res)
		if err != nil {
			return errors.New("解析结果错误"), ""
		}
		fmt.Println("返回的数据", result)
		code := result.Code
		if code == "10000" {
			detail := result.AccountDetail
			//for _, value := range detail {
			//	//tagname, _ := value.(map[string]interface{})
			//	fmt.Println("aaaa", value["account_id"])
			//}
			fmt.Println("xf", detail)
			return err, result.AccountTotalAmount
		}
		return err, ""
	}
}

//修改商户号
func UpdateMercId(id uint, mid string) {
	var c app.AppSignContract
	err := global.GVA_DB.Model(&c).Where("id = ?", id).Update("merc_id", mid).Error
	if err != nil {
		fmt.Println("修改失败原因", err.Error())
	} else {
		fmt.Println("修改成功")
	}
}

//去掉返回结果json数据中的 pre标签
func BalanceQueryTrimPrefix(data string) (err error, s BalanceQueryResult) {
	err = json.Unmarshal([]byte(data), &s)
	if err != nil {
		fmt.Println("解析错误的信息: %v", err)
	}
	return err, s
}
