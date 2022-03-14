package baozheng

import (
	"strings"
)

var _ = WithoutSpecialSymbol

// WithoutSpecialSymbol 不含特殊字符的有效名字 并且一个中文算1个字符 返回字符含有多少字符,是否有效  两边不算有效字符,中间空格算有效字符
func WithoutSpecialSymbol(str string) (l int, b bool) {
	trimStr := strings.TrimSpace(str)
	l = len([]rune(str))
	if len(str) != len(trimStr) {
		return
	}
	for _, r := range []rune(str) {
		i := int(r)
		if i < 32 || (i > 126 && i < 256) {
			return
		}
	}
	return l, true
}
