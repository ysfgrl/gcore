package ghelper

import (
	"github.com/go-playground/validator/v10"
	"github.com/ysfgrl/gcore/gerror"
)

// TODO Add custom validation
var myValidator = validator.New()

type Error struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
	Tag   string      `json:"tag"`
	Param string      `json:"param"`
}

func init() {
	myValidator.RegisterValidation("requiredBool", validateBool)
	myValidator.RegisterValidation("requiredPass", validatePass)
}
func validateBool(fl validator.FieldLevel) bool {
	return fl.Field().Bool() == true || fl.Field().Bool() == false
}
func validatePass(fl validator.FieldLevel) bool {
	val := fl.Field().String()
	if val == "" {
		return true
	}
	return false
}

func Validate(schema interface{}) *gerror.Error {
	err := myValidator.Struct(schema)
	if err != nil {
		var vErrors []Error
		for _, err := range err.(validator.ValidationErrors) {
			var el Error
			el.Field = err.Field()
			el.Value = err.Value()
			el.Tag = err.Tag()
			el.Param = err.Param()
			vErrors = append(vErrors, el)
		}

		return &gerror.Error{
			Function: "Add",
			File:     "Controller",
			Detail:   vErrors[0].Field + " " + vErrors[0].Tag,
			Code:     "api.validation_error",
			Level:    gerror.LevelError,
		}
	}
	return nil
}
