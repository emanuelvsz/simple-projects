package domain

import (
	"fmt"
	"learn/src/core/errors"
	"learn/src/utils/validator"
	"strings"

	"github.com/google/uuid"
)

var (
	IDAluno           = "O ID do aluno"
	IDAlunoInvalid    = IDAluno + "é inválido"
	NomeAluno         = "O nome do aluno"
	NomeAlunoInvalid  = NomeAluno + "é inválido"
	IdadeAluno        = "A idade do aluno"
	IdadeAlunoInvalid = IdadeAluno + "é inválida"
)

type builder struct {
	aluno         *Aluno
	invalidFields []errors.InvalidField
}

func NewBuilder() *builder {
	return &builder{aluno: &Aluno{}}
}

func (instance *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        IDAluno,
			Description: IDAlunoInvalid,
		})
		return instance
	}
	instance.aluno.id = id
	return instance
}

func (instance *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        NomeAluno,
			Description: NomeAlunoInvalid,
		})
	}

	return instance
}

func (instance *builder) WithAge(age uint8) *builder {
	if age < 18 {
		instance.invalidFields = append(instance.invalidFields, errors.InvalidField{
			Name:        IdadeAluno,
			Description: IdadeAlunoInvalid,
		})
	}
	return instance
}

func (instance *builder) Build() (*Aluno, error) {
	if len(instance.invalidFields) > 0 {
		fmt.Println("Eita deu erro")
		return nil, nil
	}
	return instance.aluno, nil
}
