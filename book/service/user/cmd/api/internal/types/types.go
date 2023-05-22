// Code generated by goctl. DO NOT EDIT.
package types

type RegisterReq struct {
	Number   string `json:"number"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type RegisterReply struct {
	RegisterResult bool   `json:"registerResult"`
	Number         string `json:"number"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}
