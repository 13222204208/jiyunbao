package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppAgreementService struct {
}

// CreateAppAgreement 创建AppAgreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAgreementService *AppAgreementService) CreateAppAgreement(appAgreement app.AppAgreement) (err error) {
	err = global.GVA_DB.Create(&appAgreement).Error
	return err
}

// DeleteAppAgreement 删除AppAgreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAgreementService *AppAgreementService) DeleteAppAgreement(appAgreement app.AppAgreement) (err error) {
	err = global.GVA_DB.Delete(&appAgreement).Error
	return err
}

// DeleteAppAgreementByIds 批量删除AppAgreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAgreementService *AppAgreementService) DeleteAppAgreementByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppAgreement{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppAgreement 更新AppAgreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAgreementService *AppAgreementService) UpdateAppAgreement(appAgreement app.AppAgreement) (err error) {
	err = global.GVA_DB.Save(&appAgreement).Error
	return err
}

// GetAppAgreement 根据id获取AppAgreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAgreementService *AppAgreementService) GetAppAgreement(id uint) (appAgreement app.AppAgreement, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appAgreement).Error
	return
}

// GetAppAgreementInfoList 分页获取AppAgreement记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAgreementService *AppAgreementService) GetAppAgreementInfoList(info appReq.AppAgreementSearch) (list []app.AppAgreement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppAgreement{})
	var appAgreements []app.AppAgreement
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Way != nil {
		db = db.Where("way = ?", info.Way)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appAgreements).Error
	return appAgreements, total, err
}

//获取协议信息，用户协议 隐私政策
func (appAgreementService *AppAgreementService) Info(info app.AppAgreement) (agreement app.AppAgreement, err error) {
	db := global.GVA_DB.Model(&app.AppAgreement{})

	if info.Way != nil {
		db = db.Where("way = ?", info.Way)
	}
	err = db.First(&agreement).Error
	return agreement, err
}
