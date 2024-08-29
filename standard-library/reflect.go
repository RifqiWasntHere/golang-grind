package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

// Generic Reflect
func getField(value any) {
	var fieldType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type : ", fieldType.Name())
	for i := 0; i < fieldType.NumField(); i++ {
		var field reflect.StructField = fieldType.Field(i)
		fmt.Println(field.Name, "Field's type is :", field.Type)
	}
}

// StructTag

type Person2 struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"30"`
	Email   string `required:"true" max:"25"`
}

func validateField(field any) {
	var fieldType reflect.Type = reflect.TypeOf(field)
	for i := 0; i < fieldType.NumField(); i++ {
		var fieldStruct reflect.StructField = fieldType.Field(i)
		fmt.Println(fieldStruct.Name, "type is :", fieldStruct.Type)
		fmt.Println(fieldStruct.Tag.Get("required"))
		fmt.Println(fieldStruct.Tag.Get("max"))
	}
	return
}

// Data Validation Example Using Struct Tag

func validatePayload(payload any) (result bool) {
	result = true
	t := reflect.TypeOf(payload)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(payload).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	// person := Person{"Rifqi"}
	// getField(person)

	person2 := Person2{"Rifqi", "Tangerang", "cihuy"}
	// validateField(person2)

	fmt.Println(validatePayload(person2))
}
