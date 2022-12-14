// 自动生成模板AppAreaInfo
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// AppAreaInfo 结构体
type AppAreaInfo struct {
      global.GVA_MODEL
      CityCd  string `json:"cityCd" form:"cityCd" gorm:"column:city_cd;comment:地区码;"`
      AreaCode  string `json:"areaCode" form:"areaCode" gorm:"column:area_code;comment:行政地区码;"`
      IsAble  string `json:"isAble" form:"isAble" gorm:"column:is_able;comment:是否可用,;"`
      CityNm  string `json:"cityNm" form:"cityNm" gorm:"column:city_nm;comment:地区码名称;"`
      AreaLevel  string `json:"areaLevel" form:"areaLevel" gorm:"column:area_level;comment:级别,1-省；2-市；3-区;"`
      ParentCityCd  string `json:"parentCityCd" form:"parentCityCd" gorm:"column:parent_city_cd;comment:父级地区代码;"`
}


// TableName AppAreaInfo 表名
func (AppAreaInfo) TableName() string {
  return "app_area_info"
}

