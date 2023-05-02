package domain

import "github.com/google/uuid"

type Person struct {
	id        uuid.UUID
	surname   string
	number    int8
	short     bool
	shirt     bool
	shirtSize string
}

func (instance Person) ID() *uuid.UUID {
	return &instance.id
}

func (instance Person) Surname() string {
	return instance.surname
}

func (instance Person) Number() int8 {
	return instance.number
}

func (instance Person) Short() bool {
	return instance.short
}

func (instance Person) Shirt() bool {
	return instance.shirt
}

func (instance Person) ShirtSize() string {
	return instance.shirtSize
}

func NewPerson(id uuid.UUID, surname, shirtSize string, number int8, short, shirt bool) *Person {
	return &Person{
		id:        id,
		surname:   surname,
		number:    number,
		short:     short,
		shirt:     shirt,
		shirtSize: shirtSize,
	}
}
