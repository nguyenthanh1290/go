// Package allergies provides functions that help determining allergy.
package allergies

var allergenDb = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Allergies returns a full list of allergies based on input score with the tesed database.
func Allergies(score uint) (allergens []string) {
	for k, v := range allergenDb {
		if score&v != 0 {
			allergens = append(allergens, k)
		}
	}
	return
}

// AllergicTo determines whether the person is allergic to the given allergen with the given score.
func AllergicTo(score uint, allergen string) bool {
	if v, prs := allergenDb[allergen]; prs {
		if score&v != 0 {
			return true
		}
	}
	return false
}
