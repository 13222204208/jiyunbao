package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppFeedbackService struct {
}

// CreateAppFeedback 创建AppFeedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFeedbackService *AppFeedbackService) CreateAppFeedback(appFeedback app.AppFeedback) (err error) {
	err = global.GVA_DB.Create(&appFeedback).Error
	return err
}

// DeleteAppFeedback 删除AppFeedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFeedbackService *AppFeedbackService) DeleteAppFeedback(appFeedback app.AppFeedback) (err error) {
	err = global.GVA_DB.Delete(&appFeedback).Error
	return err
}

// DeleteAppFeedbackByIds 批量删除AppFeedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFeedbackService *AppFeedbackService) DeleteAppFeedbackByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppFeedback{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppFeedback 更新AppFeedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFeedbackService *AppFeedbackService) UpdateAppFeedback(appFeedback app.AppFeedback) (err error) {
	err = global.GVA_DB.Save(&appFeedback).Error
	return err
}

// GetAppFeedback 根据id获取AppFeedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFeedbackService *AppFeedbackService) GetAppFeedback(id uint) (appFeedback app.AppFeedback, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appFeedback).Error
	return
}

// GetAppFeedbackInfoList 分页获取AppFeedback记录
// Author [piexlmax](https://github.com/piexlmax)
func (appFeedbackService *AppFeedbackService) GetAppFeedbackInfoList(info appReq.AppFeedbackSearch) (list []app.AppFeedback, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppFeedback{})
	var appFeedbacks []app.AppFeedback
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appFeedbacks).Error
	return appFeedbacks, total, err
}

//提交反馈内容
func (appFeedbackService *AppFeedbackService) Create(appFeedback app.AppFeedback) (err error) {
	err = global.GVA_DB.Create(&appFeedback).Error
	return err
}
