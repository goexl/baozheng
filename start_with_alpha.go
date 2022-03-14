package baozheng

import (
	`unicode`
)

var _ = StartWithAlpha

// StartWithAlpha 字符串必须以字母开头
func StartWithAlpha(str string) (checked bool) {
	if "" == str {
		return
	}
	checked = unicode.IsLetter(rune(str[0]))

	return
}
