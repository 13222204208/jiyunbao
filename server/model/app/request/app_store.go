package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AppStoreSearch struct {
	app.AppStore
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type ManageSearch struct {
	Date      string `json:"date" form:"date"`
	StartDate string `json:"startDate" form:"startDate"`
	EndDate   string `json:"endDate" form:"endDate"`
}

//财务对账
type ReconcileSearch struct {
	Month string `json:"month" form:"month"`
}

//我的老铁
type BrothersSearch struct {
	Date string `json:"date" form:"date"`
}
