package app

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	qrcode "github.com/skip2/go-qrcode"
)

type AppSetService struct {
}

// CreateAppSet 创建AppSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSetService *AppSetService) CreateAppSet(appSet app.AppSet) (err error) {
	err = global.GVA_DB.Create(&appSet).Error
	return err
}

// DeleteAppSet 删除AppSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSetService *AppSetService) DeleteAppSet(appSet app.AppSet) (err error) {
	err = global.GVA_DB.Delete(&appSet).Error
	return err
}

// DeleteAppSetByIds 批量删除AppSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSetService *AppSetService) DeleteAppSetByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppSet{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppSet 更新AppSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSetService *AppSetService) UpdateAppSet(appSet app.AppSet) (err error) {
	err = global.GVA_DB.Save(&appSet).Error
	return err
}

// GetAppSet 根据id获取AppSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSetService *AppSetService) GetAppSet(id uint) (appSet app.AppSet, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appSet).Error
	return
}

// GetAppSetInfoList 分页获取AppSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (appSetService *AppSetService) GetAppSetInfoList(info appReq.AppSetSearch) (list []app.AppSet, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppSet{})
	var appSets []app.AppSet
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appSets).Error
	return appSets, total, err
}

//url生成二维码
func (appSetService *AppSetService) Qrcode(url string) (err error, imgUrl string) {
	err, imgUrl = CreateQrcode(url)
	return err, imgUrl
}

func GetAppSet() (appSet app.AppSet) {

	err := global.GVA_DB.First(&appSet).Error
	if err != nil {
		fmt.Println("获取失败", err)
		return appSet
	} else {
		return appSet
	}
}

//获取第三方接口地址
//way 1 资料上送 2 资料修改 3图片上送，4资料确认
func GetSetUrl(way int) (err error, url string) {
	var appUrl app.AppUrl
	err = global.GVA_DB.Where("way = ?", way).First(&appUrl).Error
	url = appUrl.Url
	return
}

//生成 二维码
func CreateQrcode(qrcodeUrl string) (err error, imgUrl string) {
	qr, err := qrcode.New(qrcodeUrl, qrcode.Medium)

	imgName := utils.UserNum() + ".png"
	imgUrl = "/www/wwwroot/jiyunbao.vvv5g.com/api/uploads/qrcode/" + imgName

	imgUrl = imgUrl[36:]

	//EncryptByPub([]byte("12321323232"))

	if err != nil {
		return err, ""
	} else {
		qr.WriteFile(256, imgUrl)
		return nil, imgUrl
	}
}
