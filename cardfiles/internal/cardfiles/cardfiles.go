package cardfiles

import (
	"cardfiles/internal/person"
	"fmt"
	"github.com/cnf/structhash"
)

type Cardfile struct {
	people map[string]person.Person
}

func New(persons []person.Person)  *Cardfile {
	var people1 map[string]person.Person = map[string]person.Person{}
	for _, person:= range persons {
		hash, err := structhash.Hash(person,1)
		if err != nil {
			fmt.Errorf("An error occurred while Reading the file: open : %v", err)
			return &Cardfile{}
		}
		people1[hash] = person
	}
	return &Cardfile{
		people:  people1,
	}
}

func ChangeName (person2 *person.Person, name string) {
	person2.ChangeName(name)
}

func ChangeSurname (person2 *person.Person, surname string) {
	person2.ChangeSurname(surname)
}

func ChangePatronymic (person2 *person.Person, patronymic string) {
	person2.ChangePatronymic(patronymic)
}

