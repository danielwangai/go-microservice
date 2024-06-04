package svc

import (
	"github.com/danielwangai/twiga-foods/user-service/internal/literals"
)

func validateRegisterUserInputs(u *UserServiceRequestType) literals.Error {
	errs := literals.Error{}

	if u.FirstName == "" {
		errs["firstName"] = UserFirstNameRequiredError.Error()
	}
	if u.LastName == "" {
		errs["lastName"] = UserLastNameRequiredError.Error()
	}
	if u.Email == "" {
		errs["email"] = UserEmailRequiredError.Error()
	}
	if u.Password == "" {
		errs["password"] = UserPasswordRequiredError.Error()
	}

	return errs
}
