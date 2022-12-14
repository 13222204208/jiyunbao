package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AppPayService struct {
}

// CreateAppPay 创建AppPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPayService *AppPayService) CreateAppPay(appPay app.AppPay) (err error) {
	err = global.GVA_DB.Create(&appPay).Error
	return err
}

// DeleteAppPay 删除AppPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPayService *AppPayService) DeleteAppPay(appPay app.AppPay) (err error) {
	err = global.GVA_DB.Delete(&appPay).Error
	return err
}

// DeleteAppPayByIds 批量删除AppPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPayService *AppPayService) DeleteAppPayByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppPay{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppPay 更新AppPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPayService *AppPayService) UpdateAppPay(appPay app.AppPay) (err error) {
	err = global.GVA_DB.Save(&appPay).Error
	return err
}

// GetAppPay 根据id获取AppPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPayService *AppPayService) GetAppPay(id uint) (appPay app.AppPay, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appPay).Error
	return
}

// GetAppPayInfoList 分页获取AppPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appPayService *AppPayService) GetAppPayInfoList(info appReq.AppPaySearch) (list []app.AppPay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppPay{})
	var appPays []app.AppPay
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appPays).Error
	return appPays, total, err
}

//支付认证
func (appPayService *AppPayService) Attestation(appPay app.AppPay) (err error) {
	appPay.Status = 1
	err, _ = VerifyAppPayAttestation(appPay.Uid)
	if err == nil {
		return errors.New("已经提交了支付认证")
	}

	err = global.GVA_DB.Create(&appPay).Error

	if err == nil {
		//修改用户支付认证状态为审核中
		if err = UpdateUserState(PAY_STATE, appPay.Uid, USER_ONE); err != nil {
			return errors.New("用户支付认证状态修改失败")
		}
	}

	return err
}

//支付认证编辑
func (appPayService *AppPayService) Edit(uid uint, appPay app.AppPay) (err error) {
	appPay.Status = STATE_ONE
	err = global.GVA_DB.Model(&appPay).Where("uid = ?", uid).Updates(appPay).Error
	return err
}

//验证用户是否已经提交支付认证
func VerifyAppPayAttestation(uid uint) (err error, appPay app.AppPay) {
	if !errors.Is(global.GVA_DB.Where("uid = ?", uid).First(&appPay).Error, gorm.ErrRecordNotFound) {
		return nil, appPay
	} else {
		return errors.New("没有查询到认证记录"), appPay
	}
}
