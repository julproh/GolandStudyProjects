package person

// Структура, описывающая данные человека.
type Person struct {
	Name       string  //имя
	Surname    string  //фамилия
	Patronymic string  //отчество
	Sex        string  // пол
	Age        int     // возраст
	Height     int     //рост в см
	Weight     float32 // вес в кг
	Married    bool    //женат/замужем
}

// New - конструктор объекта типа Person.
func New(name string, surname string, patronymic string,
	sex string, age int, height int, weight float32, married bool) *Person {
	return &Person{
		Name:       name,
		Surname:    surname,
		Patronymic: patronymic,
		Sex:        sex,
		Age:        age,
		Height:     height,
		Weight:     weight,
		Married:    married}
}

// Функция, которая меняет имя (Name) человека.
func (p *Person) ChangeName(name string) {
	p.Name = name
}

// Функция, которая меняет фамилию (Surname) человека.
func (p *Person) ChangeSurname(surname string) {
	p.Surname = surname
}

// Функция, которая меняет отчество (Patronymic) человека.
func (p *Person) ChangePatronymic(patronymic string) {
	p.Patronymic = patronymic
}

// Функция, которая меняет пол (Sex) человека.
func (p *Person) ChangeSex(sex string) {
	p.Sex = sex
}

// Функция, которая меняет возраст (Age) человека.
func (p *Person) ChangeAge(age int) {
	p.Age = age
}

// Функция, которая меняет рост (Height) человека.
func (p *Person) ChangeHeight(height int) {
	p.Height = height
}

// Функция, которая меняет вес (Weight) человека.
func (p *Person) ChangeWeight(weight float32) {
	p.Weight = weight
}

// Функция, которая меняет семейное положение (Married) человека.
func (p *Person) ChangeMarried(married bool) {
	p.Married = married
}
