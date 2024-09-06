package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/victoraugustogfavaro/crud-go/src/configuration/rest_err"
)

// variáveis globais
var (
	Validate = validator.New()
	transl ut.Translator
)

// antes de rodar o main
func init() {
	// caso dê tudo certo o ok vira true
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {

		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
		// configurando a tradução
	}
}

// rest_err -> pacote 
// RestErr -> objeto
func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {
	// unmarshal, tipagem
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	// tipo inválido, chama a função passando parâmetros
	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type!")
	} else if errors.As(validation_err, &jsonValidationError) {
		// chama o array do objeto de erro
		errorsCauses := []rest_err.Causes{}

		// pra cada erro pega a mensagem e o campo e preenche no array
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			// adiciona na lista
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
