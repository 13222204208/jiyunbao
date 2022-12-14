package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppNoticeService struct {
}

// CreateAppNotice 创建AppNotice记录
// Author [piexlmax](https://github.com/piexlmax)
func (appNoticeService *AppNoticeService) CreateAppNotice(appNotice app.AppNotice) (err error) {
	err = global.GVA_DB.Create(&appNotice).Error
	return err
}

// DeleteAppNotice 删除AppNotice记录
// Author [piexlmax](https://github.com/piexlmax)
func (appNoticeService *AppNoticeService) DeleteAppNotice(appNotice app.AppNotice) (err error) {
	err = global.GVA_DB.Delete(&appNotice).Error
	return err
}

// DeleteAppNoticeByIds 批量删除AppNotice记录
// Author [piexlmax](https://github.com/piexlmax)
func (appNoticeService *AppNoticeService) DeleteAppNoticeByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppNotice{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppNotice 更新AppNotice记录
// Author [piexlmax](https://github.com/piexlmax)
func (appNoticeService *AppNoticeService) UpdateAppNotice(appNotice app.AppNotice) (err error) {
	err = global.GVA_DB.Save(&appNotice).Error
	return err
}

// GetAppNotice 根据id获取AppNotice记录
// Author [piexlmax](https://github.com/piexlmax)
func (appNoticeService *AppNoticeService) GetAppNotice(id uint) (appNotice app.AppNotice, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appNotice).Error
	return
}

// GetAppNoticeInfoList 分页获取AppNotice记录
// Author [piexlmax](https://github.com/piexlmax)
func (appNoticeService *AppNoticeService) GetAppNoticeInfoList(info appReq.AppNoticeSearch) (list []app.AppNotice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppNotice{})
	var appNotices []app.AppNotice
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appNotices).Error
	return appNotices, total, err
}

//获取公告的内容
func (appNoticeService *AppNoticeService) Info(notice app.AppNotice) (appNotice app.AppNotice, err error) {
	err = global.GVA_DB.Where("way = ?", notice.Way).First(&appNotice).Error
	return
}
