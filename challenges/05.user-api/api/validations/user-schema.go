package validations

import "fmt"

type ErrUserSchema struct {
	Message string
	Fields  map[string]string
}

func (v ErrUserSchema) Error() string {
	return fmt.Sprintf("%s: %v", v.Message, v.Fields)
}

type UserSchemaBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

func UserSchemaValidateNamesAndBio(body UserSchemaBody) error {
	firstNameLen := len(body.FirstName)
	lastNameLen := len(body.LastName)
	bioLen := len(body.Bio)

	validationSchema := map[string]string{}

	if firstNameLen < 2 || firstNameLen > 20 {
		validationSchema["first_name"] = "Must be between 2 and 20 characters"
	}
	if lastNameLen < 2 || firstNameLen > 20 {
		validationSchema["last_name"] = "Must be between 2 and 20 characters"
	}
	if bioLen < 20 || firstNameLen > 450 {
		validationSchema["biography"] = "Must be between 20 and 450 characters"
	}

	if len(validationSchema) > 0 {
		return ErrUserSchema{
			Message: "Error on validation fields",
			Fields:  validationSchema,
		}
	}

	return nil
}
