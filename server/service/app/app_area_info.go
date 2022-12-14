package app

import (
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"strconv"
)

type AppAreaInfoService struct {
}

// CreateAppAreaInfo 创建AppAreaInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAreaInfoService *AppAreaInfoService) CreateAppAreaInfo(appAreaInfo app.AppAreaInfo) (err error) {

	err, url := GetSetUrl(6)
	if err != nil {
		return errors.New("获取上送接口失败")
	}

	var i int

	for i = 0; i < 65; i++ {

		info := map[string]string{
			"pageNumber": strconv.Itoa(i),
			"pageSize":   "50",
		}
		err, res := YsPost(url, info)
		if err != nil {
			return err
		} else {
			//err, result := TrimPrefix(res)
			//if err != nil {
			//	return errors.New("解析结果错误")
			//}
			fmt.Println("返回的结果", res)
			jsons, err := simplejson.NewJson([]byte(res))
			if err != nil {
				fmt.Println("错误信息", err)
			} else {
				array, err := jsons.Array()

				for _, value := range array {

					err = global.GVA_DB.Model(&appAreaInfo).Create(value).Error
				}
				fmt.Println(len(array), err)
				//err = global.GVA_DB.Model(&appAreaInfo).Create(array).Error

			}

		}
	}
	return err

}

// DeleteAppAreaInfo 删除AppAreaInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAreaInfoService *AppAreaInfoService) DeleteAppAreaInfo(appAreaInfo app.AppAreaInfo) (err error) {
	err = global.GVA_DB.Delete(&appAreaInfo).Error
	return err
}

// DeleteAppAreaInfoByIds 批量删除AppAreaInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAreaInfoService *AppAreaInfoService) DeleteAppAreaInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppAreaInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppAreaInfo 更新AppAreaInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAreaInfoService *AppAreaInfoService) UpdateAppAreaInfo(appAreaInfo app.AppAreaInfo) (err error) {
	err = global.GVA_DB.Save(&appAreaInfo).Error
	return err
}

// GetAppAreaInfo 根据id获取AppAreaInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAreaInfoService *AppAreaInfoService) GetAppAreaInfo(id uint) (appAreaInfo app.AppAreaInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appAreaInfo).Error
	return
}

// GetAppAreaInfoInfoList 分页获取AppAreaInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (appAreaInfoService *AppAreaInfoService) GetAppAreaInfoInfoList(info appReq.AppAreaInfoSearch) (list []app.AppAreaInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppAreaInfo{})
	var appAreaInfos []app.AppAreaInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CityCd != "" {
		db = db.Where("city_cd = ?", info.CityCd)
	}
	if info.AreaCode != "" {
		db = db.Where("area_code = ?", info.AreaCode)
	}
	if info.CityNm != "" {
		db = db.Where("city_nm = ?", info.CityNm)
	}
	if info.ParentCityCd != "" {
		db = db.Where("parent_city_cd = ?", info.ParentCityCd)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appAreaInfos).Error
	return appAreaInfos, total, err
}
