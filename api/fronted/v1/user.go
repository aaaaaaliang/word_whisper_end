package v1

import "github.com/gogf/gf/v2/frame/g"

type SignInReq struct {
	g.Meta   `path:"/api/front/user/login" method:"post" tags:"用户模块"  summary:"用户登陆"`
	Username string `v:"required|length:6,16" json:"username"`
	Password string `v:"required|length:6,16" json:"password"`
}

type SignInRes struct {
	Token  string `json:"token"`
	UserID int    `json:"userId"` // 用户ID
}

type SignUpReq struct {
	g.Meta   `path:"/api/front/user/register" method:"post" tags:"用户模块"  summary:"用户注册"`
	Username string `v:"required|length:6,16" json:"username"`
	Password string `v:"required|length:6,16" json:"password"`
}

type SignUpRes struct {
	UserID int `json:"user_id"`
}
