package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AppOrderPayService struct {
}

// CreateAppOrderPay 创建AppOrderPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOrderPayService *AppOrderPayService) CreateAppOrderPay(appOrderPay app.AppOrderPay) (err error) {
	err = global.GVA_DB.Create(&appOrderPay).Error
	return err
}

// DeleteAppOrderPay 删除AppOrderPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOrderPayService *AppOrderPayService) DeleteAppOrderPay(appOrderPay app.AppOrderPay) (err error) {
	err = global.GVA_DB.Delete(&appOrderPay).Error
	return err
}

// DeleteAppOrderPayByIds 批量删除AppOrderPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOrderPayService *AppOrderPayService) DeleteAppOrderPayByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppOrderPay{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppOrderPay 更新AppOrderPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOrderPayService *AppOrderPayService) UpdateAppOrderPay(appOrderPay app.AppOrderPay) (err error) {
	err = global.GVA_DB.Save(&appOrderPay).Error
	return err
}

// GetAppOrderPay 根据id获取AppOrderPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOrderPayService *AppOrderPayService) GetAppOrderPay(id uint) (appOrderPay app.AppOrderPay, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appOrderPay).Error
	return
}

// GetAppOrderPayInfoList 分页获取AppOrderPay记录
// Author [piexlmax](https://github.com/piexlmax)
func (appOrderPayService *AppOrderPayService) GetAppOrderPayInfoList(info appReq.AppOrderPaySearch) (list []app.AppOrderPay, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppOrderPay{})
	var appOrderPays []app.AppOrderPay
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.MercId != "" {
		db = db.Where("merc_id = ?", info.MercId)
	}
	if info.PayType != 0 {
		db = db.Where("pay_type = ?", info.PayType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appOrderPays).Error
	return appOrderPays, total, err
}

//聚合支付
func (appOrderPayService *AppOrderPayService) Pay(o appReq.OrderPay) (err error, r interface{}) {
	payType := o.PayType
	fmt.Println("聚合支付提交的参数", o)
	var pay app.AppOrderPay
	pay.PayType = payType
	pay.TotalAmount = o.TotalAmount
	pay.MercId = o.MercId
	pay.Subject = o.Subject
	if payType == 1 {
		openId := HttpGet(BaseUrl + "test.php?code=" + o.Code)
		if openId == "" {
			return errors.New("获取openid失败"), r
		}
		//openId := "otXeY52vg2BQpYAwf4-73mPj_1pY"
		pay.OrderNo = utils.OrderNo() + "w"
		pay.BuyerId = openId
		global.GVA_DB.Save(&pay)
		err, r = wechatPay(pay.OrderNo, openId, o.MercId, o.TotalAmount)
	}

	if payType == 2 {
		buyerId := HttpGet(BaseUrl + "ali.php?auth_code=" + o.Code)
		fmt.Println("购买者的id", buyerId)

		if len(buyerId) != 16 {
			return errors.New("无法获取正确的user_id"), r
		}
		//buyerId := "2088612248790326"
		pay.OrderNo = utils.OrderNo() + "a"
		pay.BuyerId = buyerId
		global.GVA_DB.Save(&pay)
		err, r = aliPay(pay.OrderNo, buyerId, o.MercId, o.Subject, o.TotalAmount)

	}
	return err, r
}

//聚合支付回调
func (appOrderPayService *AppOrderPayService) Notify(orderNo string) (err error) {
	var p app.AppOrderPay
	err = global.GVA_DB.Model(&p).Where("order_no = ?", orderNo).Update("pay_status", 1).Error
	return err
}

//微信支付
func wechatPay(orderNo string, openId string, mercId string, totalAmount float64) (err error, r JsapiPayInfo) {
	t := strconv.FormatFloat(totalAmount, 'f', -1, 32)

	info := map[string]string{
		"orderNo":     orderNo,
		"mercId":      mercId,
		"openid":      openId,
		"totalAmount": t,
	}
	fmt.Println("提交的数据", info)
	url := BaseUrl + "sdk/pay/wechatPay.php"
	err, res := YsPost(url, info)
	if err != nil {
		return
	} else {
		fmt.Println("返回的", res)
		err, result := wechatPayResult(res)
		if err != nil {
			return errors.New("解析结果错误"), r
		}
		fmt.Println("返回的结果", result)
		if result.NonceStr != "" {
			return err, result
		} else {
			return errors.New("交易失败"), r
		}
	}
}

//支付宝支付
func aliPay(orderNo string, buyerId string, mercId string, subject string, totalAmount float64) (err error, r AliPayInfoResult) {
	t := strconv.FormatFloat(totalAmount, 'f', -1, 32)

	info := map[string]string{
		"orderNo":     orderNo,
		"mercId":      mercId,
		"buyerId":     buyerId,
		"totalAmount": t,
		"subject":     subject,
	}
	fmt.Println("提交的数据", info)
	url := BaseUrl + "sdk/pay/alijsapiPay.php"
	err, res := YsPost(url, info)
	if err != nil {
		return
	} else {
		fmt.Println("返回的", res)
		err, r := aliPayResult(res)
		if err != nil {
			return errors.New("解析结果错误"), r
		}
		fmt.Println("返回的结果", r)
		return err, r
	}
}

func wechatPayResult(s string) (err error, info JsapiPayInfo) {
	var r WechatPayResult
	err = json.Unmarshal([]byte(s), &r)
	if err != nil {
		fmt.Println("解析错误的信息", err)
	}
	if r.Code == "10000" && r.Msg == "Success" {
		err = json.Unmarshal([]byte(r.JsapiPayInfo), &info)
		info.OrderNo = r.OutTradeNo
		if err != nil {
			fmt.Println("解析错误", err)
		}
	}
	return err, info
}

func aliPayResult(s string) (err error, info AliPayInfoResult) {
	var r AliPayResult
	err = json.Unmarshal([]byte(s), &r)
	if err != nil {
		fmt.Println("解析错误的信息", err)
	}
	fmt.Println("解析正确的信息", r)
	if r.Code == "10000" && r.Msg == "Success" {
		err = json.Unmarshal([]byte(r.JsapiPayInfo), &info)
		info.OrderNo = r.OutTradeNo
	}
	return err, info
}

func HttpGet(url string) (res string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	res = string(body)
	fmt.Println("返回的结果", res)
	return
}
