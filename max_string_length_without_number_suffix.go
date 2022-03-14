package baozheng

import (
	`regexp`
	`unicode/utf8`
)

var _ = MaxStringLenWithoutNumberSuffix

// MaxStringLenWithoutNumberSuffix 检查去掉'-数字'尾字符串后的字符串长度是否小于等于指定长度 并且一个中文算1个字符
func MaxStringLenWithoutNumberSuffix(checkString string, maxLen int) bool {
	suffixRegexStr := `-[\d]+$`
	suffixRegex, _ := regexp.Compile(suffixRegexStr)

	matchLoc := suffixRegex.FindStringIndex(checkString)
	if 0 == len(matchLoc) {
		return utf8.RuneCountInString(checkString) <= maxLen
	}

	return utf8.RuneCountInString(checkString[:matchLoc[0]]) <= maxLen
}
