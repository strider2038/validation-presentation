package main

import (
	"context"
	"fmt"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/validator"
)

type Document struct {
	Title    string
	Keywords []string
}

func main() {
	// START OMIT
	document := Document{
		Title:    "",                                      // HL00
		Keywords: []string{"", "book", "fantasy", "book"}, // HL00
	}

	err := validator.Validate(context.Background(),
		validation.StringProperty("title", document.Title, it.IsNotBlank()),                                 // HL01
		validation.CountableProperty("keywords", len(document.Keywords), it.HasCountBetween(5, 10)),         // HL02
		validation.ComparablesProperty[string]("keywords", document.Keywords, it.HasUniqueValues[string]()), // HL03
		validation.EachStringProperty("keywords", document.Keywords, it.IsNotBlank()),                       // HL04
	)

	if violations, ok := validation.UnwrapViolationList(err); ok { // HL05
		violations.ForEach(func(i int, violation validation.Violation) error { // HL05
			fmt.Println(violation)
			return nil
		})
	}
	// Output:
	// violation at 'title': This value should not be blank.
	// violation at 'keywords': This collection should contain 5 elements or more.
	// violation at 'keywords': This collection should contain only unique elements.
	// violation at 'keywords[0]': This value should not be blank.
	// END OMIT
}
