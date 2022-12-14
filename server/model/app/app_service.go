// 自动生成模板AppService
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppService 结构体
type AppService struct {
	global.GVA_MODEL
	Title string  `json:"title" form:"title" gorm:"column:title;comment:服务名称;"`
	Price float64 `json:"price" form:"price" gorm:"column:price;comment:价格;type:decimal(10,2);default:0;"`
	Way   *int    `json:"way" form:"way" gorm:"column:way;comment:服务类型;"`
}

// TableName AppService 表名
func (AppService) TableName() string {
	return "app_service"
}
