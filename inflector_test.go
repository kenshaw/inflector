package inflector

import (
	"testing"
)

func TestPluralize(t *testing.T) {
	for _, test := range pluralTests() {
		t.Run(test.s, func(t *testing.T) {
			if s := Pluralize(test.s); s != test.exp {
				t.Errorf("Pluralize(%s) = %s, expected: %s", test.s, s, test.exp)
			}
			// Second retrieval should returns the same result. This is also tests
			// the cache
			if s := Pluralize(test.s); s != test.exp {
				t.Errorf("Pluralize(%s) = %s, expected: %s (2nd)", test.s, s, test.exp)
			}
		})
	}
}

func TestSingularize(t *testing.T) {
	for _, test := range singularTests() {
		t.Run(test.s, func(t *testing.T) {
			if s := Singularize(test.s); s != test.exp {
				t.Errorf("Singularize(%s) = %s, expected: %s", test.s, s, test.exp)
			}
			// Second retrieval should returns the same result. This is also tests
			// the cache
			if s := Singularize(test.s); s != test.exp {
				t.Errorf("Singularize(%s) = %s, expected: %s (2nd)", test.s, s, test.exp)
			}
		})
	}
}

func TestPluralizeInverse(t *testing.T) {
	for _, test := range pluralTests() {
		t.Run(test.s, func(t *testing.T) {
			if !test.match {
				return
			}
			if s := Singularize(Pluralize(test.s)); s != test.s {
				t.Errorf("Singularize(Pluralize(%s)) != %s, got: %s", test.s, test.s, s)
			}
		})
	}
}

func TestSingularizeInverse(t *testing.T) {
	for _, test := range singularTests() {
		t.Run(test.s, func(t *testing.T) {
			if !test.match {
				return
			}
			if s := Pluralize(Singularize(test.s)); s != test.s {
				t.Errorf("Pluralize(Singularize(%s)) != %s, got: %s", test.s, test.s, s)
			}
		})
	}
}

// inflectorTest is an inflector test.
type inflectorTest struct {
	s     string
	exp   string
	match bool
}

func pluralTests() []inflectorTest {
	return []inflectorTest{
		{"categoria", "categorias", true},
		{"house", "houses", true},
		{"powerhouse", "powerhouses", true},
		{"Bus", "Buses", true},
		{"bus", "buses", true},
		{"menu", "menus", true},
		{"news", "news", true},
		{"food_menu", "food_menus", true},
		{"Menu", "Menus", true},
		{"FoodMenu", "FoodMenus", true},
		{"quiz", "quizzes", true},
		{"matrix_row", "matrix_rows", true},
		{"matrix", "matrices", true},
		{"vertex", "vertices", true},
		{"index", "indices", true},
		{"Alias", "Aliases", true},
		{"Aliases", "Aliases", false},
		{"Media", "Media", true},
		{"NodeMedia", "NodeMedia", true},
		{"alumnus", "alumni", true},
		{"bacillus", "bacilli", true},
		{"cactus", "cacti", true},
		{"focus", "foci", true},
		{"fungus", "fungi", true},
		{"nucleus", "nuclei", true},
		{"octopus", "octopuses", true},
		{"radius", "radii", true},
		{"stimulus", "stimuli", true},
		{"syllabus", "syllabi", true},
		{"terminus", "termini", true},
		{"virus", "viri", true},
		{"person", "people", true},
		{"people", "people", false},
		{"glove", "gloves", true},
		{"crisis", "crises", true},
		{"tax", "taxes", true},
		{"wave", "waves", true},
		{"bureau", "bureaus", true},
		{"cafe", "cafes", true},
		{"roof", "roofs", true},
		{"foe", "foes", true},
		{"cookie", "cookies", true},
		{"wolf", "wolves", true},
		{"thief", "thieves", true},
		{"potato", "potatoes", true},
		{"hero", "heroes", true},
		{"buffalo", "buffalo", true},
		{"tooth", "teeth", true},
		{"goose", "geese", true},
		{"foot", "feet", true},
		{"objective", "objectives", true},
		{"specie", "species", false},
		{"species", "species", true},
		{"chive", "chives", true},
		{"", "", true},
	}
}

// singularTests returns the singular tests.
func singularTests() []inflectorTest {
	return []inflectorTest{
		{"categorias", "categoria", true},
		{"menus", "menu", true},
		{"news", "news", true},
		{"food_menus", "food_menu", true},
		{"Menus", "Menu", true},
		{"FoodMenus", "FoodMenu", true},
		{"houses", "house", true},
		{"powerhouses", "powerhouse", true},
		{"quizzes", "quiz", true},
		{"Buses", "Bus", true},
		{"buses", "bus", true},
		{"matrix_rows", "matrix_row", true},
		{"matrices", "matrix", true},
		{"vertices", "vertex", true},
		{"indices", "index", true},
		{"Aliases", "Alias", true},
		{"Alias", "Alias", false},
		{"Media", "Media", true},
		{"NodeMedia", "NodeMedia", true},
		{"alumni", "alumnus", true},
		{"bacilli", "bacillus", true},
		{"cacti", "cactus", true},
		{"foci", "focus", true},
		{"fungi", "fungus", true},
		{"nuclei", "nucleus", true},
		{"octopuses", "octopus", true},
		{"radii", "radius", true},
		{"stimuli", "stimulus", true},
		{"syllabi", "syllabus", true},
		{"termini", "terminus", true},
		{"viri", "virus", true},
		{"people", "person", true},
		{"gloves", "glove", true},
		{"doves", "dove", true},
		{"lives", "life", true},
		{"knives", "knife", true},
		{"wolves", "wolf", true},
		{"slaves", "slave", true},
		{"shelves", "shelf", true},
		{"taxis", "taxi", true},
		{"taxes", "tax", true},
		{"Taxes", "Tax", true},
		{"AwesomeTaxes", "AwesomeTax", true},
		{"faxes", "fax", true},
		{"waxes", "wax", true},
		{"niches", "niche", true},
		{"waves", "wave", true},
		{"bureaus", "bureau", true},
		{"genetic_analyses", "genetic_analysis", true},
		{"doctor_diagnoses", "doctor_diagnosis", true},
		{"parantheses", "paranthesis", true},
		{"Causes", "Cause", true},
		{"colossuses", "colossus", true},
		{"diagnoses", "diagnosis", true},
		{"bases", "basis", true},
		{"analyses", "analysis", true},
		{"curves", "curve", true},
		{"cafes", "cafe", true},
		{"roofs", "roof", true},
		{"foes", "foe", true},
		{"databases", "database", true},
		{"cookies", "cookie", true},
		{"thieves", "thief", true},
		{"potatoes", "potato", true},
		{"heroes", "hero", true},
		{"buffalos", "buffalo", false},
		{"babies", "baby", true},
		{"teeth", "tooth", true},
		{"geese", "goose", true},
		{"feet", "foot", true},
		{"objectives", "objective", true},
		{"species", "species", true},
		{"", "", true},
	}
}
