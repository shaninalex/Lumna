package validators

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func Validate(s any) error {
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}
