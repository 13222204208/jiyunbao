package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
)

type IndexResponse struct {
	Name        string  `json:"name" `
	State       int     `json:"state"`
	TodayAmount float64 `json:"todayAmount" default:"0.00"`
	TodayOrder  int64   `json:"todayOrder"`
}

//财务对帐
type ReconcileResponse struct {
	TodayAmount float64     `json:"todayAmount"`
	Seven       [7]float64  `json:"seven"`
	Month       interface{} `json:"month"`
}

//历史账单
type HistoryResponse struct {
	Months []*Month `json:"months"`
}

type Month struct {
	Date  string  `json:"date"`
	Money float64 `json:"money"`
}

type HistoryMonth struct {
	Days []*Day `json:"days"`
}

type Day struct {
	Date  string  `json:"date"`
	Money float64 `json:"money"`
}

//帐单详情
type BillDetails struct {
	Cycles     string    `json:"cycles"`     //结算周期
	Time       string    `json:"time"`       //结算时间
	DealTotal  float64   `json:"dealTotal"`  //交易总金额
	Order      int       `json:"order"`      //订单笔数
	Services   float64   `json:"services"`   //总服务费
	CloseTotal float64   `json:"closeTotal"` //结算总金额
	Details    []*Detail `json:"details"`    //详情明细
}

//帐单详情明细
type Detail struct {
	Nickname   string  `json:"nickname"`   //老铁昵称
	Remark     string  `json:"remark"`     //备注
	Date       string  `json:"date"`       //时间
	DealMoney  float64 `json:"dealMoney"`  //收单金额
	CloseMoney float64 `json:"closeMoney"` //结算金额
	Type       string  `json:"type"`       //线下收款
}

//24小时收数据
type IncomeResponse struct {
	Total     [24]float64 `json:"total"`
	Member    [24]float64 `json:"member"`
	Nonmember [24]float64 `json:"nonmember"`
	Records   interface{} `json:"records"`
}

//我的老铁
type BrotherResponse struct {
	Seven    [7]float64 `json:"seven"`    //近七日账单
	Brothers []*Brother `json:"brothers"` //老铁
}

type Brother struct {
	ID       uint   `json:"ID"`
	Nickname string `json:"nickname"`
	Area     string `json:"area"`
	Date     string `json:"date"`  //接入日期
	State    int    `json:"state"` //接入状态
}

//type IncomeRecords struct {
//	OrderNo  string  `json:"orderNo"`
//	Time     string  `json:"time"`
//	UserType string  `json:"userType"` //老铁，非会员
//	Amount   float64 `json:"amount"`   //增加的金额
//	Deduct   float64 `json:"deduct"`   //扣除的金额
//}

//经营数据
type ManageResponse struct {
	Turnover   float64   `json:"turnover" default:"0.00"`
	Income     float64   `json:"income" default:"0.00"`
	Rate       string    `json:"rate" default:"0%"`
	Conversion string    `json:"conversion" default:"0%"`
	Member     Member    `json:"member"`
	Nonmember  Nonmember `json:"nonmember"`
}

type Member struct {
	Order   int64   `json:"order" default:"0"`
	Amount  float64 `json:"amount" default:"0.00"`
	Deduct  float64 `json:"deduct" default:"0.00"`
	Account float64 `json:"account" default:"0.00"`
}

type Nonmember struct {
	Order   int64   `json:"order" default:"0"`
	Amount  float64 `json:"amount" default:"0.00"`
	Deduct  float64 `json:"deduct" default:"0.00"`
	Account float64 `json:"account" default:"0.00"`
}

//根据用户的ID查出mercId
func QueryMercId(uid uint) (mercId string) {
	var cust app.AppCustInfo
	global.GVA_DB.Where("uid = ?", uid).First(&cust)
	if cust.CustId == "" {
		return
	} else {
		var s app.AppSignContract
		global.GVA_DB.Where("cust_id = ?", cust.CustId).First(&s)
		if s.MercId == "" {
			return
		} else {
			return s.MercId
		}
	}
}
