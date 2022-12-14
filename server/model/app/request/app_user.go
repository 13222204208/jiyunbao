package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type AppUserSearch struct {
	app.AppUser
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type Phone struct {
	Phone string `json:"phone" form:"phone"`
}

type Login struct {
	Phone    string `json:"phone" form:"phone"`
	Code     string `json:"code" form:"code"`
	Password string `json:"password" form:"password"`
}

type UpdatePassword struct {
	OldPassword string `json:"oldPassword" form:"oldPassword"`
	NewPassword string `json:"newPassword" form:"newPassword"`
}

//忘记密码
type ForgetPassword struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
	Code     string `json:"code" form:"code"`
}

type UpdateNickname struct {
	Nickname string `json:"nickname" form:"nickname"`
}

type UpdateAvatar struct {
	Avatar string `json:"avatar" form:"avatar"`
}

//修改手机号
type UpdatePhone struct {
	OldPhone string `json:"oldPhone" form:"oldPhone"`
	OldCode  string `json:"oldCode" form:"oldCode"`
	NewPhone string `json:"newPhone" form:"newPhone"`
	NewCode  string `json:"newCode" form:"newCode"`
}
