package models

//存放前端发送请求参数的结构体

// ParamsLogin 登录接收的参数 form为接收的名称，required为必填字段
type ParamsLogin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type ParamsRegister struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required"`
	PhoneNumber int64  `json:"phone_number" form:"phone_number" binding:"required"`
}
