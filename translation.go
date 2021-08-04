package validatorx

import (
	ut `github.com/go-playground/universal-translator`
	`github.com/go-playground/validator/v10`
)

func initTranslation(validate *validator.Validate, chinese ut.Translator) (err error) {
	if err = validate.RegisterTranslation(
		"max_len_without_number_suffix",
		chinese,
		func(ut ut.Translator) error {
			return ut.Add("max_len_without_number_suffix", "{0}去掉后缀后长度必须小于等于{1}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("max_len_without_number_suffix", fe.Field(), fe.Param())
			return t
		},
	); nil != err {
		return
	}

	if err = validate.RegisterTranslation(
		"password",
		chinese,
		func(ut ut.Translator) error {
			return ut.Add("password", "密码只能由大小写字母、数字，特殊字符组成，且每种类型至少一位", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("password", fe.Field(), fe.Param())
			return t
		},
	); nil != err {
		return
	}

	return
}
