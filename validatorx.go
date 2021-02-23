package validatorx

import (
	`github.com/go-playground/locales/en`
	`github.com/go-playground/locales/zh`
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
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
	translator = ut.New(en.New(), en.New(), zh.New())
	if english, success := translator.GetTranslator("en"); success {
		if err := enLang.RegisterDefaultTranslations(validate, english); nil != err {
			return
		}
	}
	if chinese, success := translator.GetTranslator("zh"); success {
		if err := zhLang.RegisterDefaultTranslations(validate, chinese); nil != err {
			return
		}
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
	if err := initValidation(); nil != err {
		panic(err)
	}
}
