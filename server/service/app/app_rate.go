package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppRateService struct {
}

// CreateAppRate 创建AppRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRateService *AppRateService) CreateAppRate(appRate app.AppRate) (err error) {
	err = global.GVA_DB.Create(&appRate).Error
	return err
}

// DeleteAppRate 删除AppRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRateService *AppRateService) DeleteAppRate(appRate app.AppRate) (err error) {
	err = global.GVA_DB.Delete(&appRate).Error
	return err
}

// DeleteAppRateByIds 批量删除AppRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRateService *AppRateService) DeleteAppRateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppRate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppRate 更新AppRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRateService *AppRateService) UpdateAppRate(appRate app.AppRate) (err error) {
	err = global.GVA_DB.Save(&appRate).Error
	return err
}

// GetAppRate 根据id获取AppRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRateService *AppRateService) GetAppRate(id uint) (appRate app.AppRate, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appRate).Error
	return
}

// GetAppRateInfoList 分页获取AppRate记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRateService *AppRateService) GetAppRateInfoList(info appReq.AppRateSearch) (list []app.AppRate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppRate{})
	var appRates []app.AppRate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appRates).Error
	return appRates, total, err
}

func (appRateService *AppRateService) List() (list []app.AppRate, err error) {

	db := global.GVA_DB.Model(&app.AppRate{})
	var appRates []app.AppRate

	err = db.Find(&appRates).Error
	return appRates, err
}
