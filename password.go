package validatorx

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// CheckPassword 检查密码
// 密码规则
// 支持数字、字母、特殊字符
// 至少两种类型的组合
// 密码长度至少8位
func CheckPassword(password string, minLength, maxLength int) bool {
	regular := fmt.Sprintf(`/(?![\d]+$)(?![a-zA-Z]+$)(?![-=+_.,]+$)[\da-zA-Z-=+_.,]{%d,%d}/`, minLength, maxLength)
	reg := regexp.MustCompile(regular)

	return reg.MatchString(password)
}

func checkPassword(fl validator.FieldLevel) bool {
	return CheckPassword(fl.Field().String())
}
