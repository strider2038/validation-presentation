package main

import (
	"context"
	"fmt"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/validator"
)

func main() {
	banners := []struct {
		Name      string
		Keywords  []string
		Companies []string
		Brands    []string
	}{
		{Name: "Acme banner", Companies: []string{"Acme"}},
		{Name: "Empty banner"},
	}

	for _, banner := range banners {
		// START OMIT
		err := validator.Validate(
			context.Background(),
			validation.AtLeastOneOf(
				validation.CountableProperty("keywords", len(banner.Keywords), it.IsNotBlank()),
				validation.CountableProperty("companies", len(banner.Companies), it.IsNotBlank()),
				validation.CountableProperty("brands", len(banner.Brands), it.IsNotBlank()),
			),
		)
		// END OMIT
		if violations, ok := validation.UnwrapViolationList(err); ok {
			fmt.Println("banner", banner.Name, "is not valid:")
			for violation := violations.First(); violation != nil; violation = violation.Next() {
				fmt.Println(violation)
			}
		}
	}

}
