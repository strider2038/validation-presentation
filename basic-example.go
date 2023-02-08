package main

// START OMIT
import (
	"context"
	"fmt"

	"github.com/muonsoft/validation"
	"github.com/muonsoft/validation/it"
)

func main() {
	validator, _ := validation.NewValidator()

	err := validator.Validate(context.Background(),
		// список опций-аргументов
		validation.String("", // валидируемое значение // HL
			// далее список правил
			it.IsNotBlank(),      // HL
			it.HasMaxLength(100), // HL
		),
	)

	fmt.Println(err)
	// Output:
	// violation: This value should not be blank.
}

// END OMIT
