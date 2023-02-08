package main

import (
	"errors"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Document struct {
	Title    string
	Keywords []string
}

func main() {
	// START OMIT
	document := Document{
		Title:    "",
		Keywords: []string{"", "book", "fantasy", "book"},
	}

	err := validation.ValidateStruct(&document,
		validation.Field(&document.Title, validation.Required),
		validation.Field(&document.Keywords, validation.Required, validation.Length(5, 10)),
	)

	var errs validation.Errors // map[string]error
	if errors.As(err, &errs) {
		for field, e := range errs {
			fmt.Println(field+":", e.Error())
		}
	}
	// END OMIT
}
