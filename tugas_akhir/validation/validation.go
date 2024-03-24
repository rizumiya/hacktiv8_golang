package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

type CustomValidation struct {
	Validator *validator.Validate
}

func ValidationFirst(i interface{}) map[string]string {
	message := map[string]string{}
	val := CustomValidation{validator.New()}

	if err := val.Validator.Struct(i); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			switch e.Tag() {
			case "required":
				message[strings.ToLower(e.Field())] = fmt.Sprintf("Field %s is required!", e.Field())
			case "email":
				message[strings.ToLower(e.Field())] = "The input email is invalid!"
			case "min":
				message[strings.ToLower(e.Field())] = "The minimum password length is 6 characters!"
			}
		}
		return message
	}
	return nil
}
