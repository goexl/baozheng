package validatorx

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enLang "github.com/go-playground/validator/v10/translations/en"
	zhLang "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"regexp"
	"strings"
)

var translator *ut.UniversalTranslator

type Validator struct {
	validator *validator.Validate
}

// 检查手机号
func CheckMobile(mobile string) bool {
	regular := `^[+]86[-]1([3,4,5,6,7,8,9][0-9])\d{8}$`
	reg := regexp.MustCompile(regular)

	return reg.MatchString(mobile)
}

// 自定义手机号验证函数
func checkMobile(fl validator.FieldLevel) bool {
	return CheckMobile(fl.Field().String())
}

func (cv *Validator) Validate(i interface{}) (err error) {
	return cv.validator.Struct(i)
}

func New() *Validator {
	var (
		v   = validator.New()
		err error
	)
	if err = v.RegisterValidation("mobile", checkMobile); nil != err {
		log.Panicf("RegisterValidation err=%v", err.Error())
	}

	translator = ut.New(en.New(), en.New(), zh.New())
	if english, success := translator.GetTranslator("en"); success {
		if err = enLang.RegisterDefaultTranslations(v, english); nil != err {
			log.Panicf("RegisterDefaultTranslations err=%v", err.Error())
		}
	}
	if chinese, success := translator.GetTranslator("zh"); success {
		if err = zhLang.RegisterDefaultTranslations(v, chinese); nil != err {
			log.Panicf("RegisterDefaultTranslations err=%v", err.Error())
		}
	}

	return &Validator{validator: v}
}

func I18n(lang string, errs validator.ValidationErrors) (i18n validator.ValidationErrorsTranslations) {
	sep := "_"
	if strings.Contains(lang, "-") {
		sep = "-"
	}

	splits := strings.Split(lang, sep)
	for i := 0; i < len(splits); i++ {
		if t, s := translator.FindTranslator(lang); s {
			i18n = errs.Translate(t)
			break
		} else {
			if index := strings.LastIndex(lang, sep); -1 == index {
				break
			} else {
				lang = lang[0:index]
			}
		}
	}
	if nil == i18n {
		if t, s := translator.GetTranslator("zh"); s {
			i18n = errs.Translate(t)
		}
	}

	return
}
