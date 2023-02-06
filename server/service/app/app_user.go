package app

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	appRes "github.com/flipped-aurora/gin-vue-admin/server/model/app/response"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	"unsafe"
)

type AppUserService struct {
}

// CreateAppUser 创建AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) CreateAppUser(appUser app.AppUser) (err error) {
	err = global.GVA_DB.Create(&appUser).Error
	return err
}

// DeleteAppUser 删除AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteAppUser(appUser app.AppUser) (err error) {
	err = global.GVA_DB.Delete(&appUser).Error
	return err
}

// DeleteAppUserByIds 批量删除AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) DeleteAppUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppUser{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppUser 更新AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) UpdateAppUser(appUser app.AppUser) (err error) {
	err = global.GVA_DB.Save(&appUser).Error
	return err
}

// GetAppUser 根据id获取AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetAppUser(id uint) (appUser app.AppUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appUser).Error
	return
}

// GetAppUserInfoList 分页获取AppUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUserService *AppUserService) GetAppUserInfoList(info appReq.AppUserSearch) (list []app.AppUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppUser{})
	var appUsers []app.AppUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.ServiceId != nil {
		db = db.Where("service_id = ?", info.ServiceId)
	}
	if info.JoinCode != "" {
		db = db.Where("join_code LIKE ?", "%"+info.JoinCode+"%")
	}
	if info.MchType != nil {
		db = db.Where("mch_type = ?", info.MchType)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appUsers).Error
	return appUsers, total, err
}

func (appUserService *AppUserService) Register(appUser app.AppUser) (err error) {

	err, _ = VerifyAppUserIsSet(appUser.Phone)
	if err == nil {
		return errors.New("用户已经存在")
	}

	appUser.Password = utils.HashAndSalt([]byte(appUser.Password))
	err = VerifyPhoneCode(appUser.Phone, appUser.Code)
	if err != nil {
		return errors.New(err.Error())
	}
	appUser.UserNum = "A" + utils.UserNum()
	err = global.GVA_DB.Create(&appUser).Error
	if err != nil {
		return errors.New("注册失败")
	} else {
		return err
	}
}

type Info struct {
	UserInfo app.AppUser `json:"userInfo"`
	Token    string      `json:"token"`
}

//查询机构列表
func (appUserService *AppUserService) Institutions() (err error, list []app.AppUser) {
	err = global.GVA_DB.Where("is_institution = ? ", 1).Select("id", "nickname").Find(&list).Error
	return err, list
}

//用户登陆
func (appUserService *AppUserService) Login(phone, password, code string) (err error, info interface{}) {

	err, _ = VerifyAppUserIsSet(phone)
	if err != nil {
		return errors.New("用户不存在"), info
	}

	if password != "" || code != "" {
		if password != "" {
			err = VerifyAppUserPassword(phone, password)
			if err != nil {
				return errors.New("用户密码不正确"), info
			}
		}

		if code != "" {
			err = VerifyPhoneCode(phone, code)
			if err != nil {
				return errors.New(err.Error()), info
			}
		}

		err, user := VerifyAppUserIsSet(phone)
		if err != nil {
			return errors.New("用户不存在"), info
		} else {
			token, err := genToken(user.ID, user.Phone)
			if err != nil {
				return errors.New("生成token失败"), info
			} else {
				user.Password = ""
				var info Info
				info.Token = token
				info.UserInfo = user
				return err, info
			}
		}
	} else {
		return errors.New("缺少参数"), info
	}

}

//给手机号发送验证码
func (appUserService *AppUserService) SendCode(phone string) (err error) {
	err = SendSms(phone)
	if err != nil {
		return errors.New(err.Error())
	}
	return err
}

var ctx = context.Background()

//发送验证码
func SendSms(phoneNumber string) error {
	config := sdk.NewConfig()

	appSet := GetAppSet()
	fmt.Println("app设置信息", appSet)

	credential := credentials.NewAccessKeyCredential(appSet.PhoneKeyId, appSet.PhoneKeySecret)

	client, err := dysmsapi.NewClientWithOptions("cn-qingdao", config, credential)
	if err != nil {
		panic(err)
	}

	r := dysmsapi.CreateSendSmsRequest()

	r.SignName = appSet.PhoneSignName
	r.TemplateCode = appSet.PhoneTemplateCode
	r.PhoneNumbers = phoneNumber
	r.Scheme = "https"

	smsCode := utils.GenerateSmsCode(6)
	code := "{\"code\":" + smsCode + "}"

	r.TemplateParam = code

	response, err := client.SendSms(r)

	if err != nil {
		return errors.New("发送验证码错误")
	} else {
		global.GVA_REDIS.Set(ctx, phoneNumber, smsCode, time.Second*60*15)
		fmt.Printf("response is %#v\n", response)
		return nil
	}

}

//验证验证码是否正确
func VerifyPhoneCode(phoneNumber, phoneCode string) error {
	code, err := global.GVA_REDIS.Get(ctx, phoneNumber).Result()
	if err != nil {
		return errors.New("查询不到验证码")
	} else {
		fmt.Println("验证码", code, "存储的", phoneCode)
		if code == phoneCode {
			return nil
		} else {
			return errors.New("验证码不正确")
		}
	}
}

//查询用户是否已经存在
func VerifyAppUserIsSet(phoneNumber string) (err error, appUser app.AppUser) {

	if !errors.Is(global.GVA_DB.Where("phone = ?", phoneNumber).First(&appUser).Error, gorm.ErrRecordNotFound) {
		return nil, appUser
	} else {
		return errors.New("没有查询到用户"), appUser
	}
}

//验证用户的密码是否正确
func VerifyAppUserPassword(phone, password string) error {
	var appUser app.AppUser
	global.GVA_DB.Where("phone = ?", phone).First(&appUser)
	fmt.Println("密码", appUser.Password)
	ok := utils.ValidatePasswords(appUser.Password, password)

	if ok {
		return nil
	} else {
		return errors.New("登陆密码错误")
	}
}

func (appUserService *AppUserService) UpdatePassword(phone, oldPassword, newPassword string) error {
	err := VerifyAppUserPassword(phone, oldPassword)
	if err != nil {
		return errors.New("旧密码不正确")
	} else {
		password := utils.HashAndSalt([]byte(newPassword))

		err := global.GVA_DB.Model(&app.AppUser{}).Where("phone = ?", phone).Update("password", password).Error
		if err != nil {
			return errors.New("修改密码失败")
		} else {
			return err
		}
	}
}

func (appUserService *AppUserService) UpdateNickname(phone, nickname string) error {

	err := global.GVA_DB.Model(&app.AppUser{}).Where("phone = ?", phone).Update("nickname", nickname).Error
	if err != nil {
		return errors.New("修改昵称失败")
	} else {
		return err
	}
}

func (appUserService *AppUserService) UpdateAvatar(phone, avatar string) error {

	err := global.GVA_DB.Model(&app.AppUser{}).Where("phone = ?", phone).Update("avatar", avatar).Error
	if err != nil {
		return errors.New("修改头像失败")
	} else {
		return err
	}
}

//修改手机号
func (appUserService *AppUserService) UpdatePhone(oldPhone, oldCode, newPhone, newCode string) error {
	err := VerifyPhoneCode(oldPhone, oldCode)
	if err != nil {
		return errors.New(err.Error())
	}

	err = VerifyPhoneCode(newPhone, newCode)
	if err != nil {
		return errors.New(err.Error())
	}

	err = global.GVA_DB.Model(&app.AppUser{}).Where("phone = ?", oldPhone).Update("phone", newPhone).Error
	if err != nil {
		return errors.New("修改头像失败")
	} else {
		return err
	}
}

//获取用户的信息
func (appUserService *AppUserService) Info(uid uint) (err error, appUser app.AppUser, storeName string) {
	err = global.GVA_DB.Where("id = ?", uid).First(&appUser).Error

	var appStore app.AppStore
	global.GVA_DB.Where("uid = ?", uid).First(&appStore)
	return err, appUser, appStore.StoreName
}

// GenToken 生成JWT
func genToken(id uint, phone string) (string, error) {
	// 创建一个我们自己的声明
	c := appReq.MyClaims{
		id, // 自定义字段
		phone,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(appReq.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "jiyunbao",                                        // 签发人
		},
	}
	fmt.Println("token内容", c)
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(appReq.MySecret)
}

//修改用户商户或者认证的状态
//state 判断是属于商户认证还是支付认证
//s 用户状态数字
func UpdateUserState(state string, uid uint, s int) (err error) {
	var user app.AppUser
	if state == MCH_STATE {
		err = global.GVA_DB.Model(&user).Where("id = ?", uid).Update("mch_state", s).Error

	} else if state == PAY_STATE {
		err = global.GVA_DB.Model(&user).Where("id = ?", uid).Update("pay_state", s).Error

	}

	if err != nil {
		return errors.New("修改用户认证状态失败")
	} else {
		return err
	}
}

//用户忘记密码
func (appUserService *AppUserService) ForgetPassword(f appReq.ForgetPassword) error {
	phone := f.Phone
	code := f.Code
	password := f.Password
	err, _ := VerifyAppUserIsSet(phone)
	if err != nil {
		return errors.New(err.Error())
	}
	err = VerifyPhoneCode(phone, code)
	if err != nil {
		return errors.New(err.Error())
	} else {
		password := utils.HashAndSalt([]byte(password))

		err := global.GVA_DB.Model(&app.AppUser{}).Where("phone = ?", phone).Update("password", password).Error
		if err != nil {
			return errors.New("修改密码失败")
		} else {
			return err
		}
	}
}

const IdImageUrl = "http://apis.juhe.cn/idimage/verify"
const IdImageKey = "5b85a06c65d01c9dcd75caf8edadf713"

//用户身份证识别
func (appUserService *AppUserService) IdImage(image, side string) (error, interface{}) {
	baseImg, err := GetUrlImgBase64(image)
	if err != nil {
		return err, ""
	}

	postData := url.Values{}
	postData.Add("key", IdImageKey)
	postData.Add("image", baseImg)
	postData.Add("side", side)
	err, res := HttpPost(IdImageUrl, postData)
	if err != nil {
		return err, ""
	}
	if side == "front" {
		var s appRes.IdImageFrontResult
		err = json.Unmarshal([]byte(res), &s)
		if err != nil {
			fmt.Println("解析错误的信息: %v", err)
		}
		if s.ErrorCode == 0 {
			return err, s.Result
		} else {
			return errors.New(s.Reason), ""
		}
	} else {
		var s appRes.IdImageBackResult
		err = json.Unmarshal([]byte(res), &s)
		if err != nil {
			fmt.Println("解析错误的信息: %v", err)
		}
		if s.ErrorCode == 0 {
			return err, s.Result
		} else {
			return errors.New(s.Reason), ""
		}
	}

}

const BankCardUrl = "http://apis.juhe.cn/bankcardOcr/index"
const BankCardKey = "7f129959929b7d8e57818b24e40d422f"

//获取用户银行卡信息
func (appUserService *AppUserService) BankCard(image string) (error, interface{}) {
	baseImg, err := GetUrlImgBase64(image)
	if err != nil {
		return err, ""
	}

	postData := url.Values{}
	postData.Add("key", BankCardKey)
	postData.Add("image", baseImg)
	err, res := HttpPost(BankCardUrl, postData)
	if err != nil {
		return err, ""
	}
	var s appRes.BankCardResult
	err = json.Unmarshal([]byte(res), &s)
	if err != nil {
		fmt.Println("解析错误的信息: %v", err)
	}
	if s.ErrorCode == 0 {
		return err, s.Result
	} else {
		return errors.New(s.Reason), ""
	}
}

const BusUrl = "http://apis.juhe.cn/businessLicenseOcr/index"
const BusKey = "3a074340171d0d3ff923fcec0e6761f5"

//获取用户银行卡信息
func (appUserService *AppUserService) Bus(image string) (error, interface{}) {
	baseImg, err := GetUrlImgBase64(image)
	if err != nil {
		return err, ""
	}

	postData := url.Values{}
	postData.Add("key", BusKey)
	postData.Add("image", baseImg)
	err, res := HttpPost(BusUrl, postData)
	if err != nil {
		return err, ""
	}
	var s appRes.BusResult
	err = json.Unmarshal([]byte(res), &s)
	if err != nil {
		fmt.Printf("解析错误的信息: %v\n", err)
	}
	if s.ErrorCode == 0 {
		return err, s.Result
	} else {
		return errors.New(s.Reason), ""
	}
}

func HttpPost(url string, postData url.Values) (error, string) {

	response, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData.Encode()))
	if err != nil {
		return err, ""
	}
	fmt.Println("结果:", response)
	respBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("33333", err.Error())
		return err, ""
	}

	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println("44444", *str)
	return err, *str
}

func GetUrlImgBase64(path string) (baseImg string, err error) {
	str, _ := os.Getwd()
	path = str + "/" + path
	fmt.Println("当前路径", path)
	//获取本地文件
	file, err := os.Open(path)
	if err != nil {
		err = errors.New("获取本地图片失败")
		return
	}
	defer file.Close()
	imgByte, _ := ioutil.ReadAll(file)

	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
	//取图片类型
	//mimeType := http.DetectContentType(imgByte)
	//switch mimeType {
	//case "image/jpeg":
	//	baseImg = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte)
	//case "image/png":
	//	baseImg = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
	//}
	baseImg = base64.StdEncoding.EncodeToString(imgByte)
	return
}
