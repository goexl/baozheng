package validatorx

func initValidation() (err error) {
	if err = validate.RegisterValidation("mobile", checkMobile); nil != err {
		return
	}

	if err = validate.RegisterValidation("password", checkPassword); nil != err {
		return
	}

	if err = validate.RegisterValidation("without_special_symbol", checkValidateName); nil != err {
		return
	}

	return
}
