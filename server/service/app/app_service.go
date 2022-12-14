package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppServiceService struct {
}

// CreateAppService 创建AppService记录
// Author [piexlmax](https://github.com/piexlmax)
func (appServiceService *AppServiceService) CreateAppService(appService app.AppService) (err error) {
	err = global.GVA_DB.Create(&appService).Error
	return err
}

// DeleteAppService 删除AppService记录
// Author [piexlmax](https://github.com/piexlmax)
func (appServiceService *AppServiceService) DeleteAppService(appService app.AppService) (err error) {
	err = global.GVA_DB.Delete(&appService).Error
	return err
}

// DeleteAppServiceByIds 批量删除AppService记录
// Author [piexlmax](https://github.com/piexlmax)
func (appServiceService *AppServiceService) DeleteAppServiceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppService{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppService 更新AppService记录
// Author [piexlmax](https://github.com/piexlmax)
func (appServiceService *AppServiceService) UpdateAppService(appService app.AppService) (err error) {
	err = global.GVA_DB.Save(&appService).Error
	return err
}

// GetAppService 根据id获取AppService记录
// Author [piexlmax](https://github.com/piexlmax)
func (appServiceService *AppServiceService) GetAppService(id uint) (appService app.AppService, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appService).Error
	return
}

// GetAppServiceInfoList 分页获取AppService记录
// Author [piexlmax](https://github.com/piexlmax)
func (appServiceService *AppServiceService) GetAppServiceInfoList(info appReq.AppServiceSearch) (list []app.AppService, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppService{})
	var appServices []app.AppService
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appServices).Error
	return appServices, total, err
}

//获取所有服务设置信息
func (appServiceService *AppServiceService) List() (err error, appServices []app.AppService) {
	err = global.GVA_DB.Find(&appServices).Error
	return err, appServices
}
