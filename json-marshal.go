package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/muonsoft/language"
	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/it"
	"github.com/muonsoft/validation/message/translations/russian"
)

type Document struct {
	Title    string
	Keywords []string
}

func main() {
	// START TRANSLATION OMIT
	validator, _ := validation.NewValidator(validation.Translations(russian.Messages)) // HL01

	document := Document{
		Title:    "",
		Keywords: []string{"", "book", "fantasy", "book"},
	}

	ctx := language.WithContext(context.Background(), language.Russian) // HL01
	// START JSON OMIT
	err := validator.Validate(
		ctx,
		validation.StringProperty("title", document.Title, it.IsNotBlank()),
		validation.AtProperty(
			"keywords",
			validation.Countable(len(document.Keywords), it.HasCountBetween(5, 10)),
			validation.Comparables[string](document.Keywords, it.HasUniqueValues[string]()),
			validation.EachString(document.Keywords, it.IsNotBlank()),
		),
	)
	// END TRANSLATION OMIT

	data, _ := json.MarshalIndent(err, "", "\t") // HL02
	fmt.Println(string(data))
	// END JSON OMIT
	// [
	//        {
	//                "error": "is blank",
	//                "message": "Значение не должно быть пустым.",
	//                "propertyPath": "title"
	//        },
	//        {
	//                "error": "too few elements",
	//                "message": "Эта коллекция должна содержать 5 элементов или больше.",
	//                "propertyPath": "keywords"
	//        },
	//        {
	//                "error": "is not unique",
	//                "message": "Эта коллекция должна содержать только уникальные элементы.",
	//                "propertyPath": "keywords"
	//        },
	//        {
	//                "error": "is blank",
	//                "message": "Значение не должно быть пустым.",
	//                "propertyPath": "keywords[0]"
	//        }
	// ]
}
