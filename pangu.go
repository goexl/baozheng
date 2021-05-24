package validatorx

import (
	`github.com/storezhang/pangu`
)

func init() {
	app := pangu.New()

	if err := app.Sets(
		NewValidate,
		New,
	); nil != err {
		panic(err)
	}
}
