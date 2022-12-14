package app

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AppRelieveStoreService struct {
}

// CreateAppRelieveStore 创建AppRelieveStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRelieveStoreService *AppRelieveStoreService) CreateAppRelieveStore(appRelieveStore app.AppRelieveStore) (err error) {
	err = global.GVA_DB.Create(&appRelieveStore).Error
	return err
}

// DeleteAppRelieveStore 删除AppRelieveStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRelieveStoreService *AppRelieveStoreService) DeleteAppRelieveStore(appRelieveStore app.AppRelieveStore) (err error) {
	err = global.GVA_DB.Delete(&appRelieveStore).Error
	return err
}

// DeleteAppRelieveStoreByIds 批量删除AppRelieveStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRelieveStoreService *AppRelieveStoreService) DeleteAppRelieveStoreByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppRelieveStore{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppRelieveStore 更新AppRelieveStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRelieveStoreService *AppRelieveStoreService) UpdateAppRelieveStore(appRelieveStore app.AppRelieveStore) (err error) {
	err = global.GVA_DB.Save(&appRelieveStore).Error
	return err
}

// GetAppRelieveStore 根据id获取AppRelieveStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRelieveStoreService *AppRelieveStoreService) GetAppRelieveStore(id uint) (appRelieveStore app.AppRelieveStore, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appRelieveStore).Error
	return
}

// GetAppRelieveStoreInfoList 分页获取AppRelieveStore记录
// Author [piexlmax](https://github.com/piexlmax)
func (appRelieveStoreService *AppRelieveStoreService) GetAppRelieveStoreInfoList(info appReq.AppRelieveStoreSearch) (list []app.AppRelieveStore, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppRelieveStore{})
	var appRelieveStores []app.AppRelieveStore
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appRelieveStores).Error
	return appRelieveStores, total, err
}

//申请解除合作
func (appRelieveStoreService *AppRelieveStoreService) Apply(appRelieveStore app.AppRelieveStore) (err error) {
	if err = VerifyRelieveStoreApply(appRelieveStore.Uid); err != nil {
		return errors.New(err.Error())
	}
	fmt.Println("返回的数据", err)
	err = global.GVA_DB.Create(&appRelieveStore).Error
	return err
}

//查询是用户是否已经提交了解除合作记录
func VerifyRelieveStoreApply(uid uint) error {
	if !errors.Is(global.GVA_DB.Where("uid = ? AND status = ?", uid, 0).First(&app.AppRelieveStore{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("你已经提交了申请")
	} else {
		return nil
	}
}
