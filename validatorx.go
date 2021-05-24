package validatorx

import (
	`github.com/go-playground/locales/en`
	`github.com/go-playground/locales/zh`
	`github.com/go-playground/universal-translator`
	`github.com/go-playground/validator/v10`
	enLang `github.com/go-playground/validator/v10/translations/en`
	zhLang `github.com/go-playground/validator/v10/translations/zh`
)

var translator *ut.UniversalTranslator

type Validate struct {
	validate *validator.Validate
}

func (v *Validate) Struct(obj interface{}) error {
	return v.validate.Struct(obj)
}

func (v *Validate) Var(obj interface{}, tag string) error {
	return v.validate.Var(obj, tag)
}

// New 创建验证器
func New(validate *validator.Validate) *Validate {
	return &Validate{validate: validate}
}

// NewValidate 创建内置验证器
func NewValidate() (validate *validator.Validate, err error) {
	validate = validator.New()
	translator = ut.New(en.New(), zh.New())

	english, _ := translator.GetTranslator("en")
	if err = enLang.RegisterDefaultTranslations(validate, english); nil != err {
		return
	}
	chinese, _ := translator.GetTranslator("zh")
	if err = zhLang.RegisterDefaultTranslations(validate, chinese); nil != err {
		return
	}

	if err = initValidation(validate); nil != err {
		return
	}
	if err = initTranslation(validate, chinese); nil != err {
		return
	}

	return
}
