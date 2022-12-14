package mini

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mini"
	miniReq "github.com/flipped-aurora/gin-vue-admin/server/model/mini/request"
)

type MiniUserService struct {
}

// CreateMiniUser 创建MiniUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniUserService *MiniUserService) CreateMiniUser(miniUser mini.MiniUser) (err error) {
	err = global.GVA_DB.Create(&miniUser).Error
	return err
}

// DeleteMiniUser 删除MiniUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniUserService *MiniUserService) DeleteMiniUser(miniUser mini.MiniUser) (err error) {
	err = global.GVA_DB.Delete(&miniUser).Error
	return err
}

// DeleteMiniUserByIds 批量删除MiniUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniUserService *MiniUserService) DeleteMiniUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mini.MiniUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMiniUser 更新MiniUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniUserService *MiniUserService) UpdateMiniUser(miniUser mini.MiniUser) (err error) {
	err = global.GVA_DB.Save(&miniUser).Error
	return err
}

// GetMiniUser 根据id获取MiniUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniUserService *MiniUserService) GetMiniUser(id uint) (miniUser mini.MiniUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&miniUser).Error
	return
}

// GetMiniUserInfoList 分页获取MiniUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (miniUserService *MiniUserService) GetMiniUserInfoList(info miniReq.MiniUserSearch) (list []mini.MiniUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mini.MiniUser{})
	var miniUsers []mini.MiniUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.UserNum != "" {
		db = db.Where("user_num = ?", info.UserNum)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&miniUsers).Error
	return miniUsers, total, err
}

//微信小程序登陆
//func (miniUserService *MiniUserService) WeChat(code, avatar, nickname string) ([]byte, error) {
//    GetWechatOpenid(code)
//}

//获取微信小程序的openid
//func GetWechatOpenid(code string) (string, error) {
//
//}
