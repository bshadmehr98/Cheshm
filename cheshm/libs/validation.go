package libs

import (
	"log"
	"strings"

	"github.com/beego/beego/validation"
)

func CreateValidationError(form interface{}) (bool, map[string]string, error) {
	final := make(map[string]string)

	valid := validation.Validation{}
	b, err := valid.Valid(&form)

	if err != nil {
		return false, final, err
	}

	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			s := strings.Split(err.Key, ".")
			final[s[0]] = err.Message
		}
		return false, final, nil
	}

	return true, final, nil
}
