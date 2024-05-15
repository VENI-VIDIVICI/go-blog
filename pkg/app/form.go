package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (e ValidError) Error() string {
	return e.Message
}

func (e ValidErrors) Error() string {
	return strings.Join(e.Errors(), ",")
}
func (e ValidErrors) Errors() []string {
	var errs []string
	for _, es := range e {
		errs = append(errs, es.Error())
	}
	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil
}
