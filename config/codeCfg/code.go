package codeCfg

const (
	//公共操作码
	Success        = 0  // 请求成功
	InvalidRequest = -1 // 请求无效
	QueryDBErr     = -2 //查询数据库出错
	Unauthorized   = -3 //没有权限
	SignErr        = -4 // 请求无效

	NotFound = 404 //路由不存在

	ParamError        = 1001 //请求参数不正确
	PasswordError     = 1002 //账号密码不正确
	TokenError        = 1003 //token有误
	TokenExpired      = 1004
	TokenNotExist     = 1005
	CreateTokenError  = 1006
	RefreshTokenError = 1007
	LogoutTokenError  = 1008
)

//code对应描述
var Lang = map[int]string{

	InvalidRequest:    "invalid request.",
	Success:           "success.",
	QueryDBErr:        "query db error,please check db is running or query syntax is right?",
	Unauthorized:      "Unauthorized.",
	NotFound:          "route not exists.",
	ParamError:        "request param error.",
	SignErr:           "sign is error",
	TokenError:        "token is error",
	TokenNotExist:     "token is no exist",
	TokenExpired:      "token is expired",
	CreateTokenError:  "create token error",
	RefreshTokenError: "refresh token is error",
	LogoutTokenError:  "logout error",
	PasswordError:     "login password error",
}
