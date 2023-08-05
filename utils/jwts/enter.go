package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

// JwtPayLoad Payload upload in jwt
type JwtPayLoad struct {
	//Username string `json:"username"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"`
	UserID   uint   `json:"user_id"`
}

type LinkJwtPayLoad struct {
	InstanceId uint   `json:"nick_name"`
	TemplateId uint   `json:"template_id"`
	LinkType   string `json:"link_type"`
	Autofill   bool   `json:"autofill"`
}

var MySecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

type LinkCustomClaims struct {
	LinkJwtPayLoad
	jwt.StandardClaims
}
