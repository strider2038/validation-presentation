package main

import (
	"context"
	"fmt"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/validator"
)

type User struct {
	Email    string
	Password string
	City     string
}

// START USER OMIT
func (u User) Validate(ctx context.Context, validator *validation.Validator) error {
	return validator.Validate(
		ctx,
		validation.StringProperty(
			"email",
			u.Email,
			it.IsNotBlank().WhenGroups("registration"), // HL
			it.IsEmail().WhenGroups("registration"),    // HL
		),
		// ...
		// END USER OMIT
		validation.StringProperty(
			"password",
			u.Password,
			it.IsNotBlank().WhenGroups("registration"),
			it.HasMinLength(7).WhenGroups("registration"),
		),
		validation.StringProperty(
			"city",
			u.City,
			it.HasMinLength(2), // this constraint belongs to the default group
		),
	)
}

func main() {
	user := User{
		Email:    "invalid email",
		Password: "1234",
		City:     "Z",
	}

	// START VALIDATE OMIT
	err1 := validator.WithGroups("registration").Validate( // HL
		context.Background(),
		validation.Valid(user),
	)
	err2 := validator.Validate(
		context.Background(),
		validation.Valid(user),
	)
	// END VALIDATE OMIT

	if violations, ok := validation.UnwrapViolationList(err1); ok {
		fmt.Println("violations for registration group:")
		for violation := violations.First(); violation != nil; violation = violation.Next() {
			fmt.Println(violation)
		}
	}
	if violations, ok := validation.UnwrapViolationList(err2); ok {
		fmt.Println("violations for default group:")
		for violation := violations.First(); violation != nil; violation = violation.Next() {
			fmt.Println(violation)
		}
	}

}
