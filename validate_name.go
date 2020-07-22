package validatorx

import (
	`github.com/go-playground/validator/v10`
)

// 不含特殊字符的有效名字 并且一个中文算1个字符 返回字符含有多少字符,是否有效
func ValidateName(str string) (int, bool) {
	l := len([]rune(str))
	for _, r := range []rune(str) {
		i := int(r)

		if i <= 32 || (i >= 126 && i < 256) {
			return 0, false
		}
	}
	return l, true
}

// 不含特殊字符的有效名字 并且一个中文算1个字符
func checkValidateName(fl validator.FieldLevel) bool {
	_, valid := ValidateName(fl.Field().String())
	return valid
}
