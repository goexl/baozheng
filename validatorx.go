package validatorx

import (
	`github.com/go-playground/locales/en`
	`github.com/go-playground/locales/zh`
	`github.com/go-playground/universal-translator`
	`github.com/go-playground/validator/v10`
	enLang `github.com/go-playground/validator/v10/translations/en`
	zhLang `github.com/go-playground/validator/v10/translations/zh`
)

var (
	translator *ut.UniversalTranslator
	validate   *validator.Validate
)

type Validatorx struct {
	validator *validator.Validate
}

func (v *Validatorx) Validate(obj interface{}) error {
	return v.validator.Struct(obj)
}

// New 创建
func New() (validator *Validatorx) {
	translator = ut.New(en.New(), zh.New())
	english, _ := translator.GetTranslator("en")
	if err := enLang.RegisterDefaultTranslations(validate, english); nil != err {
		panic(err)
	}
	chinese, _ := translator.GetTranslator("zh")
	if err := zhLang.RegisterDefaultTranslations(validate, chinese); nil != err {
		panic(err)
	}

	if err := initValidation(validate); nil != err {
		panic(err)
	}
	if err := initTranslation(validate, chinese); nil != err {
		panic(err)
	}
	validator = &Validatorx{validator: validate}

	return
}

// Validate 验证
func Validate(obj interface{}) error {
	return validate.Struct(obj)
}

func init() {
	validate = validator.New()
}
