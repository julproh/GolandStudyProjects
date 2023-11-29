package cardfiles

import (
	"fmt"

	"github.com/cnf/structhash"

	"cardfiles/person"
)

type Cardfiles struct {
	people map[string]person.Person
}

// оапрвап
func New() *Cardfiles {
	return &Cardfiles{
		people: make(map[string]person.Person),
	}
}

func (cardfile Cardfiles) AddToCardfile(person person.Person) error {
	hash, err := structhash.Hash(person, 1)
	if err != nil {
		fmt.Errorf("An error occurred: %v", err)
	}
	cardfile.people[hash] = person
	return nil
}
func (cardfile Cardfiles) ChangeName(person2 *person.Person, name string) {
	person2.ChangeName(name)
}

func (cardfile Cardfiles) ChangeSurname(person2 *person.Person, surname string) {
	person2.ChangeSurname(surname)
}

func (cardfile Cardfiles) ChangePatronymic(person2 *person.Person, patronymic string) {
	person2.ChangePatronymic(patronymic)
}

func (cardfile Cardfiles) ChangeAge(person2 *person.Person, age int) {
	person2.ChangeAge(age)
}

func (cardfile Cardfiles) ChangeSex(person2 *person.Person, sex string) {
	person2.ChangeSex(sex)
}

func (cardfile Cardfiles) ChangeHeight(person2 *person.Person, height int) {
	person2.ChangeHeight(height)
}
func (cardfile Cardfiles) ChangeWeight(person2 *person.Person, weight float32) {
	person2.ChangeWeight(weight)
}

func (cardfile Cardfiles) ChangeMarried(person2 *person.Person, married bool) {
	person2.ChangeMarried(married)
}

func (cardfile Cardfiles) DeletePersoneByHash(hash string) {
	delete(cardfile.people, hash)
}

func (cardfile Cardfiles) DeletePersone(person2 *person.Person) {
	hash, err := structhash.Hash(&person2, 1)
	if err != nil {
		fmt.Errorf("An error occurred: %v", err)
	}
	delete(cardfile.people, hash)
}
