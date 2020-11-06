package serverCfg

//接口列表
const (
	GetSignAPi = "getSign" //測試sign
	SetSignAPi = "setSign" //設置sign
)

//接口列表映射
var ServerApi = map[string]map[string]string{
	GetSignAPi: {
		"api":    "/api/sign/get",
		"method": "get",
	},
	SetSignAPi: {
		"api":    "/api/sign/set",
		"method": "post",
	},
}
