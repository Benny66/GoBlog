package sign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goBlog/config/default"
	"io/ioutil"
	"goBlog/common/alarm"
	"goBlog/common/func"
	"goBlog/common/response"
	"goBlog/config/code"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"strconv"
)

func Sign() gin.HandlerFunc {
	return func(c *gin.Context) {

		sign, err := verifySign(c)
		if sign != nil {
			c.JSON(http.StatusUnauthorized, response.NewResponse(codeCfg.SignErr, sign))
			c.Abort()
			return
		}
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.NewResponse(codeCfg.SignErr, err))
			c.Abort()
			return
		}

		c.Next()
	}
}

// 验证签名
func verifySign(c *gin.Context) (res map[string]string, err error) {
	var method = c.Request.Method
	var timestamp int64
	var sign string
	var createSign string
	var req url.Values
	//var debug string
	type s map[string]interface{}
	raw := make(s)

	req = c.Request.URL.Query()
	paramsStr, _ := json.Marshal(req)
	alarm.New("request params: " + string(paramsStr))
	//当前请求的ip
	//ipAddress := c.ClientIP()
	sign = c.Query("sign")
	timestamp, _ = strconv.ParseInt(c.Query("timestamp"), 10, 64)
	//获取当前ip的key和过期时间,没有则选默认
	key := defaultCfg.Cfg.SignKey
	expire := defaultCfg.Cfg.SignExpire
	if method == "POST" || method == "PUT" {
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		json.Unmarshal(body, &raw)
		paramsStr, _ := json.Marshal(raw)
		alarm.New("body params: " + string(paramsStr))
		raw["timestamp"] = timestamp
	}
	// 验证过期时间
	now := _func.GetTimeUnix()
	if timestamp > now || now-timestamp >= expire {
		err = errors.New("Timestamp Error")
		return
	}

	// 验证签名
	if sign == "" {
		err = errors.New("sign Error")
		return
	}
	if method == "POST" || method == "PUT" {
		createSign = CreateBodySign(raw, key)
	} else {
		createSign = CreateQuerySign(req, key)
	}
	if sign != createSign {
		alarm.New("sign error,fromSign:" + createSign + " sign:" + sign)
		err = errors.New("sign Error")
		return
	}
	return
}

// 生成签名
func CreateQuerySign(params url.Values, signKey string) string {
	var key []string
	var str = ""
	for k := range params {
		if k != "sign" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	// 自定义签名算法
	sign := _func.MD5(str + "&key=" + signKey)
	return sign
}

func CreateBodySign(params map[string]interface{}, signKey string) string {
	var key []string
	var str = ""
	for k, _ := range params {
		if k != "sign" {
			key = append(key, k)
		}
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params[key[i]])
		}
	}
	// 自定义签名算法
	sign := _func.MD5(str + "&key=" + signKey)
	return sign
}
