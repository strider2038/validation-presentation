package main

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

type Product struct {
	Name      string    `json:"name" valid:"required"`
	Component Component `json:"component"`
}

type Component struct {
	ID   int    `json:"id"`
	Name string `json:"name" valid:"required"`
}

func main() {
	p := Product{
		Name:      "",
		Component: Component{ID: 1, Name: ""},
	}

	result, err := govalidator.ValidateStruct(p)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println(result)
}
