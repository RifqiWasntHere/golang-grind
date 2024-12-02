package api

import (
	"simplebank/util"

	"github.com/go-playground/validator/v10"
)

var validateCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		// Validates if currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
