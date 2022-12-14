// 自动生成模板AppFacilitatingAgency
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppFacilitatingAgency 结构体
type AppFacilitatingAgency struct {
	global.GVA_MODEL
	Phone         string `json:"phone" form:"phone" gorm:"column:phone;comment:手机号;"`
	Name          string `json:"name" form:"name" gorm:"column:name;comment:;"`
	Way           *int   `json:"way" form:"way" gorm:"column:way;comment:;default:1"`
	Principal     string `json:"principal" form:"principal" gorm:"column:principal;comment:;"`
	Card          string `json:"card" form:"card" gorm:"column:card;comment:;"`
	Mail          string `json:"mail" form:"mail" gorm:"column:mail;comment:;"`
	Area          string `json:"area" form:"area" gorm:"column:area;comment:;"`
	Address       string `json:"address" form:"address" gorm:"column:address;comment:;"`
	Status        *int   `json:"status" form:"status" gorm:"column:status;comment:;default:0"`
	Certification *int   `json:"certification" form:"certification" gorm:"column:certification;comment:;default:0"`
}

// TableName AppFacilitatingAgency 表名
func (AppFacilitatingAgency) TableName() string {
	return "app_facilitating_agency"
}
