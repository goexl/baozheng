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
	validatorx *Validatorx
)

type Validatorx struct {
	validate *validator.Validate
}

func (v *Validatorx) Validate(obj interface{}) error {
	return v.validate.Struct(obj)
}

// GetInstance 获取数据验证器
func GetInstance() *Validatorx {
	return validatorx
}

// Struct 验证结构体
func Struct(obj interface{}) error {
	return validate.Struct(obj)
}

// Var 验证
func Var(obj interface{}, tag string) error {
	return validate.Var(obj, tag)
}

func init() {
	validate = validator.New()
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

	validatorx = &Validatorx{validate: validate}
}
