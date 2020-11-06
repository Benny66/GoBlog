package users

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

// 定义 Person 结构体
type LoginParams struct {
	Username string `form:"username" json:"username" binding:"required,CheckUsername"`
	Password string `form:"password"  json:"password"  binding:"required"`
}
var CheckUsername validator.Func = func(fl validator.FieldLevel) bool {
	username, ok := fl.Field().Interface().(string)
	if ok {
		fmt.Println(username)
		matched, _ := regexp.MatchString("^[a-zA-Z0-9]{6,16}$", username)
		fmt.Println(matched)
		if matched {
			return true
		}
	}
	return false
}