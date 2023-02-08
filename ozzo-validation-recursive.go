package main

import (
	"errors"
	"fmt"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Product struct {
	Name       string
	Tags       []string
	Components []Component
}

func (p Product) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Name, validation.Required),
		validation.Field(&p.Tags, validation.Required, validation.Length(5, 0)),
		validation.Field(&p.Components, validation.Length(1, 0)),
	)
}

type Components []Component

func (cs Components) Validate() error {
	errs := make(validation.Errors, len(cs))
	for i, c := range cs {
		errs[strconv.Itoa(i)] = c.Validate()
	}
	return errs.Filter()
}

type Component struct {
	ID   int
	Name string
	Tags []string
}

func (c Component) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required),
		validation.Field(&c.Tags, validation.Length(1, 0)),
	)
}

func main() {
	p := Product{
		Name:       "",
		Tags:       []string{"device", "", "phone", "device"},
		Components: []Component{{ID: 1, Name: ""}},
	}

	err := p.Validate()

	var errs validation.Errors // map[string]error
	if errors.As(err, &errs) {
		for field, e := range errs {
			fmt.Println(field+":", e.Error())
		}
	} else {
		fmt.Println(err)
	}
}
