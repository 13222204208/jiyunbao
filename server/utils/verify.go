package utils

var (
	IdVerify               = Rules{"ID": []string{NotEmpty()}}
	ApiVerify              = Rules{"Path": {NotEmpty()}, "Description": {NotEmpty()}, "ApiGroup": {NotEmpty()}, "Method": {NotEmpty()}}
	MenuVerify             = Rules{"Path": {NotEmpty()}, "ParentId": {NotEmpty()}, "Name": {NotEmpty()}, "Component": {NotEmpty()}, "Sort": {Ge("0")}}
	MenuMetaVerify         = Rules{"Title": {NotEmpty()}}
	LoginVerify            = Rules{"CaptchaId": {NotEmpty()}, "Captcha": {NotEmpty()}, "Username": {NotEmpty()}, "Password": {NotEmpty()}}
	RegisterVerify         = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}, "AuthorityId": {NotEmpty()}}
	PageInfoVerify         = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	CustomerVerify         = Rules{"CustomerName": {NotEmpty()}, "CustomerPhoneData": {NotEmpty()}}
	AutoCodeVerify         = Rules{"Abbreviation": {NotEmpty()}, "StructName": {NotEmpty()}, "PackageName": {NotEmpty()}, "Fields": {NotEmpty()}}
	AutoPackageVerify      = Rules{"PackageName": {NotEmpty()}}
	AuthorityVerify        = Rules{"AuthorityId": {NotEmpty()}, "AuthorityName": {NotEmpty()}}
	AuthorityIdVerify      = Rules{"AuthorityId": {NotEmpty()}}
	OldAuthorityVerify     = Rules{"OldAuthorityId": {NotEmpty()}}
	ChangePasswordVerify   = Rules{"Password": {NotEmpty()}, "NewPassword": {NotEmpty()}}
	SetUserAuthorityVerify = Rules{"AuthorityId": {NotEmpty()}}

	//app用户注册
	AppUserRegister = Rules{"NickName": {NotEmpty()}, "Phone": {RegexpMatch(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`)}, "Password": {NotEmpty()}, "Code": {NotEmpty()}}

	//app用户修改密码
	AppUserUpdatePassword = Rules{"OldPassword": {NotEmpty()}, "NewPassword": {NotEmpty()}}

	//app忘记密码
	AppUserForgetPassword = Rules{"Phone": {RegexpMatch(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`)}, "Code": {NotEmpty()}, "Password": {NotEmpty()}}

	//app用户修改用户名昵称
	AppUserUpdateNickname = Rules{"Nickname": {NotEmpty()}}

	//app用户修改用户头像
	AppUserUpdateAvatar = Rules{"Avatar": {NotEmpty()}}

	//app修改手机号
	AppUserUpdatePhone = Rules{"OldPhone": {RegexpMatch(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`)}, "OldCode": {NotEmpty()}, "NewPhone": {RegexpMatch(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`)}, "NewCode": {NotEmpty()}}

	//短信发送验证手机号
	UserCode = Rules{"Phone": {RegexpMatch(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`)}}

	//商户进件中的资料上传
	CustInfoVerify = Rules{
		"MercName":       {NotEmpty()},
		"MercType":       {NotEmpty()},
		"MccCd":          {NotEmpty()},
		"ContactMail":    {NotEmpty()},
		"CusMgrNm":       {NotEmpty()},
		"CrpCertNo":      {NotEmpty()},
		"CrpCertType":    {NotEmpty()},
		"CertBgn":        {NotEmpty()},
		"CertExpire":     {NotEmpty()},
		"CrpNm":          {NotEmpty()},
		"CrpPhone":       {NotEmpty()},
		"StlAccNo":       {NotEmpty()},
		"BankSubCode":    {NotEmpty()},
		"StlAccType":     {NotEmpty()},
		"BusProviceCode": {NotEmpty()},
		"BusCityCode":    {NotEmpty()},
		"BusAddr":        {NotEmpty()},
	}

	MiniUserLoginVerify = Rules{
		"Code":     {NotEmpty()},
		"Avatar":   {NotEmpty()},
		"Nickname": {NotEmpty()},
		"MiniType": {NotEmpty()},
	}
)
