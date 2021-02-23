package validatorx

import (
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	translator *ut.UniversalTranslator
	validate   *validator.Validate
)

type Validator struct {
	validator *validator.Validate
}

func init() {
	validate = validator.New()
	if err := initValidation(); nil != err {
		panic(err)
	}
}
