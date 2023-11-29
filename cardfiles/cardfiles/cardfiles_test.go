package cardfiles

import (
	"testing"

	"cardfiles/person"
)

// Тест: Добавление информации о новом человеке в картотеку.
func TestCardfiles_AddToCardfile(t *testing.T) {
	cardFile := New()
	person1 := person.New("Ivan",
		"Ivanov",
		"Ivanovich",
		"M",
		40,
		178,
		80.2,
		true,
	)
	cardFile.AddToCardfile(*person1)
	if len(cardFile.people) != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %s.", len(cardFile.people), "1")
	}

}
