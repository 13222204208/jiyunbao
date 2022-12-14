package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mini"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    miniReq "github.com/flipped-aurora/gin-vue-admin/server/model/mini/request"
)

type MiniSetService struct {
}

// CreateMiniSet 创建MiniSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniSetService *MiniSetService) CreateMiniSet(miniSet mini.MiniSet) (err error) {
	err = global.GVA_DB.Create(&miniSet).Error
	return err
}

// DeleteMiniSet 删除MiniSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniSetService *MiniSetService)DeleteMiniSet(miniSet mini.MiniSet) (err error) {
	err = global.GVA_DB.Delete(&miniSet).Error
	return err
}

// DeleteMiniSetByIds 批量删除MiniSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniSetService *MiniSetService)DeleteMiniSetByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mini.MiniSet{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMiniSet 更新MiniSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniSetService *MiniSetService)UpdateMiniSet(miniSet mini.MiniSet) (err error) {
	err = global.GVA_DB.Save(&miniSet).Error
	return err
}

// GetMiniSet 根据id获取MiniSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniSetService *MiniSetService)GetMiniSet(id uint) (miniSet mini.MiniSet, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&miniSet).Error
	return
}

// GetMiniSetInfoList 分页获取MiniSet记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniSetService *MiniSetService)GetMiniSetInfoList(info miniReq.MiniSetSearch) (list []mini.MiniSet, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&mini.MiniSet{})
    var miniSets []mini.MiniSet
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&miniSets).Error
	return  miniSets, total, err
}
