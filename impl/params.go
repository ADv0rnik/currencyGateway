package impl

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Params struct {
	BaseUrl string `validate:"required,url"`
	City    string `validate:"required"`
}

func RetrieveParams(reqData map[string]string) (Params, error) {
	p := Params{
		BaseUrl: reqData["baseUrl"],
		City:    reqData["city"],
	}

	validate := validator.New()

	if err := validate.Struct(&p); err != nil {
		return Params{}, fmt.Errorf("retrieve implementation parameters: %s", strings.ReplaceAll(err.Error(), "\n", ", "))
	}
	return p, nil
}
