package inflector_test

import (
	"fmt"

	"github.com/kenshaw/inflector"
)

func ExampleSingularize() {
	fmt.Println(inflector.Singularize("People"))
	// Output:
	// Person
}

func ExamplePluralize() {
	fmt.Println(inflector.Pluralize("octopus"))
	// Output:
	// octopuses
}
