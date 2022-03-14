package baozheng

import (
	"strings"
)

var _ = PrefixOrSuffixSpace

// PrefixOrSuffixSpace 检查前后缀空白字符
func PrefixOrSuffixSpace(str string) bool {
	trimStr := strings.TrimSpace(str)

	return len(str) == len(trimStr)
}
