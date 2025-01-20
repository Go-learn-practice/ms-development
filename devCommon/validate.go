package common

import "regexp"

// VerifyMobile 验证手机号
func VerifyMobile(mobile string) bool {
	if mobile == "" {
		return false
	}

	regular := "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

// VerifyEmailFormat 邮箱格式验证
func VerifyEmailFormat(email string) bool {
	pattern := ``
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
