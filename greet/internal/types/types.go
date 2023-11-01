// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"` // parameters are auto validated
}

type Response struct {
	Message string `json:"message"`
}

type User struct {
	Id              uint   `json:"id"`
	Username        string `json:"username"`
	Gender          int    `json:"gender"`
	Email           string `json:"email"`
	Role            string `json:"role"`
}

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginRes struct {
	Info  User   `json:"info"`
	Token string `json:"token"`
}

type RegisterReq struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Tel       int    `form:"tel"`
	AvatarUrl string `form:"avatar_url"`
}

type RegisterRes struct {
	Uid uint32 `json:"uid"`
}