package go_validation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator"
)

func TestValidatorInit(t *testing.T) {
	validate := validator.New()

	if validate == nil {
		t.Error("validator is null")
	}
}

func TestValidateField(t *testing.T) {
	validate := validator.New()
	user := "This a string"
	err := validate.Var(user, "required") //Validate a single variable
	if err != nil {                       //There's also option to validate a context
		fmt.Println(err)
	}
}

func TestValidateMultipleField(t *testing.T) {
	validate := validator.New()
	user := "This a string"
	validValue := "This a string"
	err := validate.VarWithValue(user, validValue, "eqfield") //Validate whether value of var 1 and 2 are equal (effect of "eqfield" tag)
	if err != nil {
		fmt.Println(err)
	}
}

func TestValidateMultipleTags(t *testing.T) {
	validate := validator.New()
	payload := "897123129863"

	err := validate.Var(payload, "required,numeric") //Supports multiple tags
	if err != nil {
		fmt.Println(err)
	}
}

func TestValidateParameterTags(t *testing.T) {
	validate := validator.New()
	payload := "99999999999"

	err := validate.Var(payload, "required,numeric,min=10,max=100") //Params are activated by "=", and it is relative. Means if it
	if err != nil {
		fmt.Println(err)
	}
}

func TestValidateStruct(t *testing.T) {
	validate := validator.New()

	type ExampleStruct struct {
		Name  string `validate:"required,min=5,max=75"`
		Email string `validate:"required,email"`
	}

	payload := ExampleStruct{
		Name:  "Rifqi",
		Email: "YEhaaw@outlook.com",
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

// func TestValidationErrors(t *testing.T) {
// 	validate := validator.New()

// 	type ExampleStruct struct {
// 		Name  string `validate:"required,min=5,max=75"`
// 		Email string `validate:"required,email"`
// 	}

// 	payload := ExampleStruct{
// 		Name:  "Rifqi",
// 		Email: "YEhaawoutlook.com",
// 	}

// 	err := validate.Struct(payload)
// 	if err != nil {
// 		validationErrors := err.(validator.ValidationErrors) //Is actually an ALIAS of existing []FieldError
// 		for _, fieldError := range validationErrors {
// 			println("field error :", fieldError.Field(), "| on tag", fieldError.Tag(), " : ", fieldError)
// 		}
// 	}
// }

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "eko",
		Password: "eko",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for i, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", validationErrors[i])
		}
	}
}

// Cross-Field (Cuma Eqsfield yang dikasih value)
func TestCrossField(t *testing.T) {
	validate := validator.New()

	type ExampleStruct struct {
		Name            string `validate:"required,min=5,max=75"`
		Email           string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	payload := ExampleStruct{
		Name:            "Rifqi",
		Email:           "YEhaaw@outlook.com",
		Password:        "akulaku",
		ConfirmPassword: "akulaku",
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

func TestNestedStruct(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"` //Cukup include nested struct nya disini, tapi si nested struct itu sendiri harus di punya property validate juga
	}

	validate := validator.New()

	payload := User{
		Id:   "1",
		Name: "Iqfir",
		Address: Address{
			City:    "Tangerang",
			Country: "Indonesia",
		},
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

func TestNestedCollection(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"` //Tambah property [dive] untuk Field bertipe collection
	}

	validate := validator.New()

	payload := User{
		Id:   "1",
		Name: "Iqfir",
		Addresses: []Address{
			{
				City:    "Foo",
				Country: "Bar",
			},
			{
				City:    "Foo2",
				Country: "Bar2",
			},
		},
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

func TestBasicCollection(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`       //Tambah tag [dive] untuk Field bertipe collection
		Hobbies   []string  `validate:"required,dive,min=3"` //Khusus basic collection (non-struct). Tambahin tag untuk isi dari collectionnya setelah [dive]
	}

	validate := validator.New()

	payload := User{
		Id:   "1",
		Name: "Iqfir",
		Addresses: []Address{
			{
				City:    "Foo",
				Country: "Bar",
			},
			{
				City:    "Foo2",
				Country: "Bar2",
			},
		},
		Hobbies: []string{
			"X",
			"Mancing",
			"Badminton",
		},
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

func TestValidateMap(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id      string            `validate:"required"`
		Name    string            `validate:"required"`
		Address Address           `validate:"required"`
		Schools map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
		//Untuk map, validation bisa diaplikasikan ke key. Caranya pake dive, kemudian kasih tag keys yang ditutup dengan endkeys.
		//Dan semua tags yang ada diantara keys-endkeys itu bakal jadi tag validator untuk key.
	}

	validate := validator.New()

	payload := User{
		Id:   "1",
		Name: "Iqfir",
		Address: Address{
			City:    "Tangerang",
			Country: "Indonesia",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SDI Yakmi",
			},
			"SMP": {
				Name: "",
			},
		},
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

func TestValidateBasicMap(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id      string            `validate:"required"`
		Name    string            `validate:"required"`
		Address Address           `validate:"required"`
		Schools map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
		//Untuk map, validation bisa diaplikasikan ke key. Caranya pake dive, kemudian kasih tag keys yang ditutup dengan endkeys.
		//Dan semua tags yang ada diantara keys-endkeys itu bakal jadi tag validator untuk key. **Dive terakhir ada untuk mengakses values**
		Wallet map[string]int `validate:"dive,keys,required,endkeys,required,gt=1000"`
		//Untuk basic map (non-struct) cuma sedikit beda sama yg struct. Bedanya, value dari basic map gak perlu di dive.
		//Semua tags yang ada setelah endkeys dianggap sebagai tag validator untuk value
	}

	validate := validator.New()

	payload := User{
		Id:   "1",
		Name: "Iqfir",
		Address: Address{
			City:    "Tangerang",
			Country: "Indonesia",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SDI Yakmi",
			},
			"SMP": {
				Name: "Yeahh",
			},
		},
		Wallet: map[string]int{
			"Cash":    100000,
			"Mandiri": 9,
		},
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAliasTag(t *testing.T) {

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id      string            `validate:"required"`
		Name    string            `validate:"required"`
		Address Address           `validate:"required"`
		Schools map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
		Wallet  map[string]int    `validate:"basicmap,required,gt=1000"` //Alias tag-nya dipake disini
	}

	validate := validator.New()
	validate.RegisterAlias("basicmap", "dive,keys,required,endkeys") //Alias tag bisa berlaku untuk multiple tags, berguna banget buat bikin preset

	payload := User{
		Id:   "1",
		Name: "Iqfir",
		Address: Address{
			City:    "Tangerang",
			Country: "Indonesia",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SDI Yakmi",
			},
			"SMP": {
				Name: "Yeahh",
			},
		},
		Wallet: map[string]int{
			"1": 100000,
			"":  100000,
		},
	}

	err := validate.Struct(payload)
	if err != nil {
		fmt.Println(err)
	}
}
