package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AppMchService struct {
}

// CreateAppMch 创建AppMch记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMchService *AppMchService) CreateAppMch(appMch app.AppMch) (err error) {
	err = global.GVA_DB.Create(&appMch).Error
	return err
}

// DeleteAppMch 删除AppMch记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMchService *AppMchService) DeleteAppMch(appMch app.AppMch) (err error) {
	err = global.GVA_DB.Delete(&appMch).Error
	return err
}

// DeleteAppMchByIds 批量删除AppMch记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMchService *AppMchService) DeleteAppMchByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppMch{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppMch 更新AppMch记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMchService *AppMchService) UpdateAppMch(appMch app.AppMch) (err error) {
	err = global.GVA_DB.Save(&appMch).Error
	return err
}

// GetAppMch 根据id获取AppMch记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMchService *AppMchService) GetAppMch(id uint) (appMch app.AppMch, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appMch).Error
	return
}

// GetAppMchInfoList 分页获取AppMch记录
// Author [piexlmax](https://github.com/piexlmax)
func (appMchService *AppMchService) GetAppMchInfoList(info appReq.AppMchSearch) (list []app.AppMch, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppMch{})
	var appMchs []app.AppMch
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FirmName != "" {
		db = db.Where("firm_name LIKE ?", "%"+info.FirmName+"%")
	}
	if info.StoreName != "" {
		db = db.Where("store_name LIKE ?", "%"+info.StoreName+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appMchs).Error
	return appMchs, total, err
}

//商户认证
func (appMchService *AppMchService) Attestation(appMch app.AppMch) (err error) {

	err, _ = VerifyAppMchAttestation(appMch.Uid)
	if err == nil {
		return errors.New("商户已经提交了认证")
	}
	appMch.Status = 1
	err = global.GVA_DB.Create(&appMch).Error

	if err == nil {
		//修改用户支付认证状态为审核中
		if err = UpdateUserState(MCH_STATE, appMch.Uid, USER_ONE); err != nil {
			return errors.New("用户商户认证状态修改失败")
		}
	}

	return err
}

//修改商户认证信息
func (appMchService *AppMchService) Edit(uid uint, appMch app.AppMch) error {
	appMch.Status = STATE_ONE
	err := global.GVA_DB.Model(&appMch).Where("uid = ?", uid).Updates(appMch).Error
	return err
}

//验证用户是否已经提交商户认证
func VerifyAppMchAttestation(uid uint) (err error, appMch app.AppMch) {
	if !errors.Is(global.GVA_DB.Where("uid = ?", uid).First(&appMch).Error, gorm.ErrRecordNotFound) {
		return nil, appMch
	} else {
		return errors.New("没有查询到认证记录"), appMch
	}
}
