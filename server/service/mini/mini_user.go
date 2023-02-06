package mini

import (
	"context"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/mini"
	miniReq "github.com/flipped-aurora/gin-vue-admin/server/model/mini/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/golang-jwt/jwt/v4"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/business"
	"github.com/silenceper/wechat/v2/miniprogram/config"
	"gorm.io/gorm"
	"time"
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

//根据用户ID获取用户信息
func (miniUserService *MiniUserService) Info(id uint) (miniUser mini.MiniUser, err error) {
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

//支付宝小程序登陆
func (miniUserService *MiniUserService) AliPayLogin(code string) (err error, m mini.MiniUser) {
	err, set := GetMinAppSet(2)
	if err != nil {
		return errors.New("获取支付宝小程序配置失败"), m
	}
	rsp, err := alipay.SystemOauthToken(context.Background(), set.Appid, set.Private, "authorization_code", code, "RSA2")
	if err != nil {
		xlog.Error(err)
		return
	}
	authToken := rsp.Response.AccessToken
	fmt.Println("authToken", set.Appid)
	client, err := alipay.NewClient(set.Appid, set.Private, true)
	if err != nil {
		xlog.Error(err)
		return
	}
	info, err := client.UserInfoShare(context.Background(), authToken)
	if err != nil {
		xlog.Error(err)
		return
	}
	res := info.Response
	openid := res.UserId
	avatar := res.Avatar
	nickname := res.NickName
	fmt.Println("结果", res)
	return SaveMiniUser(openid, avatar, nickname, 2)
}

//微信小程序登陆
func (miniUserService *MiniUserService) WeChatLogin(code, avatar, nickname string) (err error, m mini.MiniUser) {
	err, openid := GetWechatOpenid(code)
	if err != nil {
		return err, m
	}
	return SaveMiniUser(openid, avatar, nickname, 1)
}

//获取支付宝小程序的手机号
func (miniUserService *MiniUserService) GetAliPayPhone(encryptedData string, uid uint) (err error) {
	phone := new(alipay.UserPhone)
	// 解密支付宝开放数据
	//    encryptedData:包括敏感数据在内的完整用户信息的加密数据
	//    secretKey:AES密钥，支付宝管理平台配置
	//    beanPtr:需要解析到的结构体指针
	secretKey := "UA/1C/yLSRzWUo0lxWMrzg=="
	err = alipay.DecryptOpenDataToStruct(encryptedData, secretKey, phone)
	fmt.Println("错误信息", err)
	xlog.Infof("%+v", phone)
	if phone.Code != "10000" {
		return errors.New(phone.Msg)
	} else {
		global.GVA_DB.Model(&mini.MiniUser{}).Where("id = ?", uid).Update("phone", phone.Mobile)
	}

	return err
}

//获取微信小程序用户的手机号
func (miniUserService *MiniUserService) GetWeChatPhone(code string, uid uint) (err error) {
	err, MinAppSet := GetMinAppSet(1)
	if err != nil {
		return errors.New("获取微信小程序配置失败")
	}
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     MinAppSet.Appid,
		AppSecret: MinAppSet.AppSecret,
		Cache:     memory,
	}
	min := wc.GetMiniProgram(cfg)
	b := min.GetBusiness()
	var p business.GetPhoneNumberRequest
	p.Code = code
	info, err := b.GetPhoneNumber(&p)
	if err != nil {
		return errors.New("获取微信小程序用户手机号失败")
	}
	phone := info.PurePhoneNumber
	global.GVA_DB.Model(&mini.MiniUser{}).Where("id = ?", uid).Update("phone", phone)
	fmt.Println("info", info, err)
	return
}

//获取用户的商户码和贵宾码
func (miniUserService *MiniUserService) GetQrcode(t int) (err error, img string) {
	if t == 1 {
		img = "uploads/qrcode/merchant.png"
		return err, img
	} else {
		img = "uploads/qrcode/vip.png"
		return err, img
	}
}

type MoneyDetail struct {
	ID      uint    `json:"ID"`
	Way     string  `json:"way"`
	Time    string  `json:"time"`
	Money   float64 `json:"money"`
	Balance float64 `json:"balance"`
}

//获取用户零钱及明细
func (miniUserService *MiniUserService) GetMoneyDetail() (err error, list []*MoneyDetail) {
	list = append(list, &MoneyDetail{
		1, "提现金额", "2023-01-31", 0.00, 0.00,
	})
	list = append(list, &MoneyDetail{
		2, "推广红包", "2023-01-31", 0.00, 0.00,
	})
	return err, list
}

// GenToken 生成JWT
func genToken(id uint, userNum string) (string, error) {
	// 创建一个我们自己的声明
	c := miniReq.MyClaims{
		id, // 自定义字段
		userNum,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(miniReq.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-mini-project",                                  // 签发人
		},
	}
	fmt.Println("token内容", c)
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(miniReq.MySecret)
}

//获取微信小程序的openid
func GetWechatOpenid(code string) (err error, s string) {
	err, MinAppSet := GetMinAppSet(1)
	if err != nil {
		return errors.New("获取微信小程序配置失败"), s
	}
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &config.Config{
		AppID:     MinAppSet.Appid,
		AppSecret: MinAppSet.AppSecret,
		Cache:     memory,
	}
	min := wc.GetMiniProgram(cfg)
	a := min.GetAuth()
	r, errs := a.Code2Session(code)

	if errs != nil {
		return errors.New(errs.Error()), s
	} else {
		return errs, r.OpenID
	}
}

func GetMinAppSet(miniType int) (err error, m mini.MiniSet) {
	err = global.GVA_DB.Where("mini_type = ?", miniType).First(&m).Error
	return err, m
}

func SaveMiniUser(openid, avatar, nickname string, miniType int) (err error, m mini.MiniUser) {
	if !errors.Is(global.GVA_DB.Where("openid = ?", openid).First(&m).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		global.GVA_DB.Model(&m).Where("openid = ?", openid).Updates(map[string]interface{}{"avatar": avatar, "nickname": nickname})
		//return errors.New("用户已注册"), m
		token, err := genToken(m.ID, m.UserNum)
		if err != nil {
			return errors.New("获取token失败"), m
		} else {
			m.Token = token
			return err, m
		}
	} else {
		m.Openid = openid
		m.Avatar = avatar
		m.Nickname = nickname
		m.UserNum = utils.UserNum()
		m.MiniType = miniType
		err = global.GVA_DB.Create(&m).Error
		if err != nil {
			return errors.New("保存失败"), m
		} else {
			token, err := genToken(m.ID, m.UserNum)
			if err != nil {
				return errors.New("获取token失败"), m
			} else {
				m.Token = token
				return err, m
			}
		}

	}
}
