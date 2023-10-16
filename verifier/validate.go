package verifier

import (
	"encoding/json"
	"hyphen-backend-hellog/cerrors"
	"hyphen-backend-hellog/cerrors/exception"

	"github.com/go-playground/validator/v10"
)

func Validate(modelValidate interface{}) {
	validate := validator.New()
	err := validate.Struct(modelValidate)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is " + err.Tag(),
			})
		}

		jsonMessage, errJson := json.Marshal(messages)
		exception.Sniff(errJson)

		panic(cerrors.ValidationError{
			ErrorMessage: string(jsonMessage),
		})
	}
}
