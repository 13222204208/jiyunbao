package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AppStoreService struct {
}

// CreateAppStore 创建AppStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appStoreService *AppStoreService) CreateAppStore(appStore app.AppStore) (err error) {
	err = global.GVA_DB.Create(&appStore).Error
	return err
}

// DeleteAppStore 删除AppStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appStoreService *AppStoreService) DeleteAppStore(appStore app.AppStore) (err error) {
	err = global.GVA_DB.Delete(&appStore).Error
	return err
}

// DeleteAppStoreByIds 批量删除AppStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appStoreService *AppStoreService) DeleteAppStoreByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppStore{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppStore 更新AppStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appStoreService *AppStoreService) UpdateAppStore(appStore app.AppStore) (err error) {
	err = global.GVA_DB.Save(&appStore).Error
	return err
}

// GetAppStore 根据id获取AppStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appStoreService *AppStoreService) GetAppStore(id uint) (appStore app.AppStore, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appStore).Error
	return
}

// GetAppStoreInfoList 分页获取AppStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appStoreService *AppStoreService) GetAppStoreInfoList(info appReq.AppStoreSearch) (list []app.AppStore, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppStore{})
	var appStores []app.AppStore
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StoreName != "" {
		db = db.Where("store_name LIKE ?", "%"+info.StoreName+"%")
	}
	if info.StorePhone != "" {
		db = db.Where("store_phone = ?", info.StorePhone)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appStores).Error
	return appStores, total, err
}

//门店设置
func (appStoreService *AppStoreService) Edit(store app.AppStore) (err error) {
	uid := store.Uid
	err, _ = VerifyAppStoreSet(uid)
	if err != nil {
		err = global.GVA_DB.Create(&store).Error
		return err
	} else {
		err = global.GVA_DB.Model(&store).Where("uid = ?", uid).Updates(store).Error
		return err
	}
}

//门店详情
func (appStoreService *AppStoreService) Detail(uid uint) (err error, store app.AppStore) {
	err = global.GVA_DB.Where("uid = ?", uid).First(&store).Error
	return
}

//验证用户是否已经提交支付认证
func VerifyAppStoreSet(uid uint) (err error, appStore app.AppStore) {
	if !errors.Is(global.GVA_DB.Where("uid = ?", uid).First(&appStore).Error, gorm.ErrRecordNotFound) {
		return nil, appStore
	} else {
		return errors.New("没有查询到门店记录"), appStore
	}
}
