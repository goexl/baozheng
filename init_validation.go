package validatorx

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func initValidation(validate *validator.Validate) (err error) {
	if err = validate.RegisterValidation("mobile", checkMobile); nil != err {
		return
	}
	if err = validate.RegisterValidation("password", checkPassword); nil != err {
		return
	}
	if err = validate.RegisterValidation("without_special_symbol", checkWithoutSpecialSymbol); nil != err {
		return
	}
	if err = validate.RegisterValidation("filename", checkValidFileNamePath); nil != err {
		return
	}
	if err = validate.RegisterValidation("start_with_alpha", checkStartWithAlpha); nil != err {
		return
	}
	if err = validate.RegisterValidation("prefix_or_suffix_space", checkPrefixOrSuffixSpace); nil != err {
		return
	}
	if err = validate.RegisterValidation("max_len_without_number_suffix", checkMaxStringLenWithoutNumberSuffix); nil != err {
		return
	}

	return
}

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

	return
}
