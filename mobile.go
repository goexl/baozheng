package baozheng

import (
	"regexp"
)

var _ = Mobile

// Mobile 检查手机号
func Mobile(mobile string) bool {
	regular := `^[+]86[-]1([3,4,5,6,7,8,9][0-9])\d{8}$`
	reg := regexp.MustCompile(regular)

	return reg.MatchString(mobile)
}
