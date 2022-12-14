package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type AppFacilitatingAgencyService struct {
}

// CreateAppFacilitatingAgency 创建AppFacilitatingAgency记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFacilitatingAgencyService *AppFacilitatingAgencyService) CreateAppFacilitatingAgency(appFacilitatingAgency app.AppFacilitatingAgency) (err error) {
	err = global.GVA_DB.Create(&appFacilitatingAgency).Error
	return err
}

// DeleteAppFacilitatingAgency 删除AppFacilitatingAgency记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFacilitatingAgencyService *AppFacilitatingAgencyService)DeleteAppFacilitatingAgency(appFacilitatingAgency app.AppFacilitatingAgency) (err error) {
	err = global.GVA_DB.Delete(&appFacilitatingAgency).Error
	return err
}

// DeleteAppFacilitatingAgencyByIds 批量删除AppFacilitatingAgency记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFacilitatingAgencyService *AppFacilitatingAgencyService)DeleteAppFacilitatingAgencyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppFacilitatingAgency{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAppFacilitatingAgency 更新AppFacilitatingAgency记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFacilitatingAgencyService *AppFacilitatingAgencyService)UpdateAppFacilitatingAgency(appFacilitatingAgency app.AppFacilitatingAgency) (err error) {
	err = global.GVA_DB.Save(&appFacilitatingAgency).Error
	return err
}

// GetAppFacilitatingAgency 根据id获取AppFacilitatingAgency记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFacilitatingAgencyService *AppFacilitatingAgencyService)GetAppFacilitatingAgency(id uint) (appFacilitatingAgency app.AppFacilitatingAgency, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appFacilitatingAgency).Error
	return
}

// GetAppFacilitatingAgencyInfoList 分页获取AppFacilitatingAgency记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFacilitatingAgencyService *AppFacilitatingAgencyService)GetAppFacilitatingAgencyInfoList(info appReq.AppFacilitatingAgencySearch) (list []app.AppFacilitatingAgency, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&app.AppFacilitatingAgency{})
    var appFacilitatingAgencys []app.AppFacilitatingAgency
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Phone != "" {
        db = db.Where("phone = ?",info.Phone)
    }
    if info.Name != "" {
        db = db.Where("name = ?",info.Name)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&appFacilitatingAgencys).Error
	return  appFacilitatingAgencys, total, err
}
