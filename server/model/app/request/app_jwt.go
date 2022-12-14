package request

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

type MyClaims struct {
    ID      uint   `json:"id"`
    Phone string `json:"phone"`
    jwt.StandardClaims
}

//然后我们定义JWT的过期时间，这里以72小时为例：
const TokenExpireDuration = time.Hour * 24 * 30 * 12

//接下来还需要定义Secret
var MySecret = []byte("token20221024")
