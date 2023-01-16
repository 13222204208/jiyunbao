package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mini"
	miniReq "github.com/flipped-aurora/gin-vue-admin/server/model/mini/request"
)

type MiniClassifyService struct {
}

// CreateMiniClassify 创建MiniClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniClassifyService *MiniClassifyService) CreateMiniClassify(miniClassify mini.MiniClassify) (err error) {
	err = global.GVA_DB.Create(&miniClassify).Error
	return err
}

// DeleteMiniClassify 删除MiniClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniClassifyService *MiniClassifyService) DeleteMiniClassify(miniClassify mini.MiniClassify) (err error) {
	err = global.GVA_DB.Delete(&miniClassify).Error
	return err
}

// DeleteMiniClassifyByIds 批量删除MiniClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniClassifyService *MiniClassifyService) DeleteMiniClassifyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mini.MiniClassify{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMiniClassify 更新MiniClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniClassifyService *MiniClassifyService) UpdateMiniClassify(miniClassify mini.MiniClassify) (err error) {
	err = global.GVA_DB.Save(&miniClassify).Error
	return err
}

// GetMiniClassify 根据id获取MiniClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniClassifyService *MiniClassifyService) GetMiniClassify(id uint) (miniClassify mini.MiniClassify, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&miniClassify).Error
	return
}

// GetMiniClassifyInfoList 分页获取MiniClassify记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniClassifyService *MiniClassifyService) GetMiniClassifyInfoList(info miniReq.MiniClassifySearch) (list []mini.MiniClassify, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mini.MiniClassify{})
	var miniClassifys []mini.MiniClassify
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&miniClassifys).Error
	return miniClassifys, total, err
}

//获取小程序首页分类
func (miniClassifyService *MiniClassifyService) List(pid string) (list []mini.MiniClassify, err error) {
	err = global.GVA_DB.Where("pid = ? AND status = ?", pid, 1).Find(&list).Error
	return
}
