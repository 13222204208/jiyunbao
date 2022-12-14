package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"go.uber.org/zap"
)

type AppUrlService struct {
}

// CreateAppUrl 创建AppUrl记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUrlService *AppUrlService) CreateAppUrl(appUrl app.AppUrl) (err error) {
	err = global.GVA_DB.Create(&appUrl).Error
	return err
}

// DeleteAppUrl 删除AppUrl记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUrlService *AppUrlService) DeleteAppUrl(appUrl app.AppUrl) (err error) {
	err = global.GVA_DB.Delete(&appUrl).Error
	return err
}

// DeleteAppUrlByIds 批量删除AppUrl记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUrlService *AppUrlService) DeleteAppUrlByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppUrl{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppUrl 更新AppUrl记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUrlService *AppUrlService) UpdateAppUrl(appUrl app.AppUrl) (err error) {
	err = global.GVA_DB.Save(&appUrl).Error
	return err
}

// GetAppUrl 根据id获取AppUrl记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUrlService *AppUrlService) GetAppUrl(id uint) (appUrl app.AppUrl, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appUrl).Error
	return
}

// GetAppUrlInfoList 分页获取AppUrl记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUrlService *AppUrlService) GetAppUrlInfoList(info appReq.AppUrlSearch) (list []app.AppUrl, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppUrl{})
	var appUrls []app.AppUrl
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appUrls).Error
	return appUrls, total, err
}

//查询第三方接口第地址
func ThirdPartyUrl(way int) string {
	var appUrl app.AppUrl
	err := global.GVA_DB.Where(" way = ?", way).First(&appUrl).Error
	if err != nil {
		global.GVA_LOG.Error("查询第三方接口失败!", zap.Error(err))
	}
	return appUrl.Url

}
