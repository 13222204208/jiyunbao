package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	appRes "github.com/flipped-aurora/gin-vue-admin/server/model/app/response"
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

		if store.StoreName != "" {
			global.GVA_DB.Model(&app.AppMch{}).Where("uid = ?", uid).Update("store_name", store.StoreName)
		}

		if store.OperatingStatus > 0 {
			global.GVA_DB.Model(&app.AppMch{}).Where("uid = ?", uid).Update("operating_status", store.OperatingStatus)
		}
		return err
	}
}

//门店详情
func (appStoreService *AppStoreService) Detail(uid uint) (err error, store app.AppStore) {
	err = global.GVA_DB.Where("uid = ?", uid).First(&store).Error
	return
}

//门店首页
func (appStoreService *AppStoreService) Index(uid uint) (err error, i appRes.IndexResponse) {
	var store app.AppMch
	err = global.GVA_DB.Where("uid = ?", uid).First(&store).Error
	i.Name = store.StoreName
	i.State = store.OperatingStatus

	mercId := appRes.QueryMercId(uid)
	if mercId != "" {

	}
	return err, i

}

//经营数据
func (appStoreService *AppStoreService) Manage(uid uint, date string) (err error, m appRes.ManageResponse) {
	mercId := appRes.QueryMercId(uid)
	if mercId != "" {

	}
	m.Rate = "0%"
	m.Conversion = "0%"
	return err, m
}

//24小时收数据
func (appStoreService *AppStoreService) Income(uid uint, date string) (err error, i appRes.IncomeResponse) {
	mercId := appRes.QueryMercId(uid)
	var r []app.AppOrderPay
	if mercId != "" {
		global.GVA_DB.Where("merc_id = ? AND pay_time LIKE ?", mercId, date+"%").Find(&r)
	}
	i.Records = r

	return err, i
}

//财务对账
func (appStoreService *AppStoreService) Reconcile(uid uint, month string) (err error, i appRes.ReconcileResponse) {
	mercId := appRes.QueryMercId(uid)
	var r []app.AppOrderPay
	if mercId != "" {
		if month != "" {
			global.GVA_DB.Where("merc_id = ? AND pay_time LIKE ?", mercId, month+"%").Find(&r)
		}
	}
	i.Month = r
	return err, i
}

//历史账单
func (appStoreService *AppStoreService) History(uid uint, month string) (err error, list interface{}) {
	if month != "" {
		var d []*appRes.Day
		d = append(d, &appRes.Day{Date: "2023-01-15", Money: 1.11})
		d = append(d, &appRes.Day{Date: "2023-01-16", Money: 0.11})
		var h appRes.HistoryMonth
		h.Days = d
		return err, h
	} else {
		var m []*appRes.Month
		m = append(m, &appRes.Month{Date: "2022-12", Money: 1.11})
		m = append(m, &appRes.Month{Date: "2023-01", Money: 2.21})
		var h appRes.HistoryResponse
		h.Months = m
		return err, h
	}
}

//账单详情
func (appStoreService *AppStoreService) Bill(uid uint) (err error, data *appRes.BillDetails) {
	var d []*appRes.Detail
	d = append(d, &appRes.Detail{
		Nickname: "老铁昵称", Remark: "扣除服务费收单手续费0.25", Date: "2023-01-15 15:04:05", DealMoney: 10.00, CloseMoney: 9.25, Type: "线下收款",
	})
	d = append(d, &appRes.Detail{
		Nickname: "老铁昵称", Remark: "扣除服务费收单手续费0.25", Date: "2023-01-15 15:04:05", DealMoney: 10.00, CloseMoney: 9.25, Type: "线下收款",
	})
	data = &appRes.BillDetails{
		Cycles:     "2022-12-01 - 2023-01-31",
		Time:       "2023-01-31 00:01:32",
		DealTotal:  100.12,
		Order:      3,
		Services:   5.01,
		CloseTotal: 96.01,
		Details:    d,
	}
	return err, data
}

//我的老铁
func (appStoreService *AppStoreService) Brothers(uid uint, date string) (err error, data appRes.BrotherResponse) {
	var b []*appRes.Brother
	b = append(b, &appRes.Brother{
		ID: 1, Nickname: "老铁昵称", Area: "贵州贵阳", Date: "2023-01-02", State: 1,
	})
	b = append(b, &appRes.Brother{
		ID: 2, Nickname: "老铁昵称", Area: "2023-01-02", Date: "2023-01-03", State: 0,
	})
	data.Brothers = b
	return err, data
}

//验证用户是否已经提交支付认证
func VerifyAppStoreSet(uid uint) (err error, appStore app.AppStore) {
	if !errors.Is(global.GVA_DB.Where("uid = ?", uid).First(&appStore).Error, gorm.ErrRecordNotFound) {
		return nil, appStore
	} else {
		return errors.New("没有查询到门店记录"), appStore
	}
}
