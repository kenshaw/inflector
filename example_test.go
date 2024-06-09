package inflector_test

import (
	"fmt"

	"github.com/kenshaw/inflector"
)

func ExampleSingularize() {
	for _, s := range []string{
		"People",
		"Archives",
		"octopuses",
	} {
		fmt.Println(inflector.Singularize(s))
	}
	// Output:
	// Person
	// Archive
	// octopus
}

func ExamplePluralize() {
	for _, s := range []string{
		"Person",
		"Archive",
		"octopus",
	} {
		fmt.Println(inflector.Pluralize(s))
	}
	// Output:
	// People
	// Archives
	// octopuses
}
