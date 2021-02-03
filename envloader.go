package envloader

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func Load(configStruct interface{}) (error, interface{}) {
	v := reflect.New(reflect.TypeOf(configStruct)).Interface()

	if err := env.Parse(v); err != nil {
		return err, nil
	}

	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		return err, nil
	}

	return nil, v
}
