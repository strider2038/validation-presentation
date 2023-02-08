package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

// START OMIT
type Document struct {
	Title   string   `valid:"required"`
	Keyword string   `valid:"stringlength(5|10)"`
	Tags    []string // нет встроенной проверки на длину слайса
}

func main() {
	document := Document{
		Title:   "",
		Keyword: "book",
	}

	result, err := govalidator.ValidateStruct(document)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println(result)
	// Output:
	// error: Keyword: book does not validate as stringlength(5|10);Title: non zero value required
	// false
}

// END OMIT
