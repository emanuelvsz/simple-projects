package domain

import "github.com/google/uuid"

type Aluno struct {
	id   uuid.UUID
	name string
	age  uint8
}

func (instance *Aluno) GetID() uuid.UUID {
	return instance.id
}

func (instance *Aluno) GetName() string {
	return instance.name
}

func (instance *Aluno) GetAge() uint8 {
	return instance.age
}
