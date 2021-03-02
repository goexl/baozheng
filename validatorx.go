package validatorx

import (
	`sync`

	`github.com/go-playground/locales/en`
	`github.com/go-playground/locales/zh`
	`github.com/go-playground/universal-translator`
	`github.com/go-playground/validator/v10`
	enLang `github.com/go-playground/validator/v10/translations/en`
	zhLang `github.com/go-playground/validator/v10/translations/zh`
)

var (
	once       sync.Once
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

// Init 初始化数据验证器
func Init() *Validatorx {
	once.Do(func() {
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
	})

	return validatorx
}

// Validate 验证
func Validate(obj interface{}) error {
	return validate.Struct(obj)
}
