package baozheng

import (
	`path/filepath`
	`regexp`
)

var (
	_ = Base
	_ = Filename
	_ = Filepath
)

// Base 有效文件名（Windows标准）/路径和文件名检测同一个
// 我们用200个字符
func Base(file string) bool {
	if file == "/" {
		return true
	}

	fileRegexStr := `^[^\\\./:\*\?\"<>\|]{1}[^\\/:\*\?\"<>\|]{0,200}$`
	filenameRegex := regexp.MustCompile(fileRegexStr)

	return filenameRegex.MatchString(filepath.Base(file))
}

// Filename 有效文件名（Windows标准）
func Filename(filename string) bool {
	fileRegexStr := `^[^\\\./:\*\?\"<>\|]{1}[^\\/:\*\?\"<>\|]{0,254}$`
	filenameRegex := regexp.MustCompile(fileRegexStr)

	return filenameRegex.MatchString(filename)
}

// Filepath 有效文件夹名
func Filepath(filepath string) bool {
	pathRegexStr := `^[^\\\/\?\*\&quot;\'\&gt;\&lt;\:\|]*$`
	pathRegex := regexp.MustCompile(pathRegexStr)
	if len(filepath) == 0 {
		return false
	}
	filepath = filepath[1:]

	return pathRegex.MatchString(filepath)
}
