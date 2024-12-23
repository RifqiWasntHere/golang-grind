package go_validation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator"
)

func ValidateUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	// Line diatas itu agak membingungkan, jadi sebenernya field.Field().Interface itu cuma nge-return value string aja.
	// Tapi type assertion .(string) itu return boolean walapun gak ada cara buat cari taunya (misal, ngehover mouse ke  .(string))

	if ok {
		if len(value) < 5 {
			return false
		}
		if value != strings.ToUpper(value) {
			return false
		}
	}
	return true
}
func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", ValidateUsername) // Function custom validation harus di regist dulu

	type LoginRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required"`
	}

	payload := LoginRequest{
		Username: "RIFQi",
		Password: "qewqweqwe",
	}
	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func ValidatePin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		fmt.Println("error getting parameter")
	}

	value := field.Field().String()
	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length //This line returns a boolean. why ? cuz it equals to if len(value) == length{ return true }
}

func TestValidatePin(t *testing.T) {
	validate := validator.New()

	validate.RegisterValidation("pin", ValidatePin)

	type Credentials struct {
		Pin string `validate:"required,pin=6"`
	}

	payload := Credentials{
		Pin: "918273",
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

// Or rule di validator
func TestOrRule(t *testing.T) {
	validate := validator.New()

	type LoginRequest struct {
		Username string `validate:"required,email|numeric"` //Just add a pipe symbol for OR condition!
		Password string `validate:"required"`
	}

	payload := LoginRequest{
		Username: "as@gmail.com",
		Password: "ahaide",
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

// Custom validation for cross-field (checking condition within 2 fields)
func EqualNotSensitive(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2() //To understand of what this method returns, read the docs!
	//reflect.Value = field's value | reflect.Kind = field's datatype | bool #1 = nullable or not | bool #2 = is the field retrieval success ?\
	if !ok {
		fmt.Println("error : field not found")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestEqualNotSensitive(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("eqns", EqualNotSensitive)
	type ChangePassword struct {
		OldPassword string `validate:"required"`
		NewPassword string `validate:"required,eqns=OldPassword"` //Just the same like using eqsfield.
		//OR rule can be applied like = `validate:"required,eqns=OldPassword|eqns=AnotherFIeld"`
	}

	payload := ChangePassword{
		OldPassword: "ikimANEH",
		NewPassword: "IKIManehh",
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}
