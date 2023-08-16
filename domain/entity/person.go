package entity

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type Person struct {
	ID        uuid.UUID `json:"id"`
	Nickname  string    `json:"apelido"`
	Name      string    `json:"nome"`
	Birthdate string    `json:"nascimento"`
	Stack     []string  `json:"stack"`
}

type PersonParams struct {
	ID        uuid.UUID `json:"id"`
	Nickname  string    `json:"apelido"`
	Name      string    `json:"nome"`
	Birthdate string    `json:"nascimento"`
	Stack     []string  `json:"stack"`
}

func NewPerson(params PersonParams) (person *Person, err error) {
	person = &Person{
		ID:        params.ID,
		Nickname:  params.Nickname,
		Name:      params.Name,
		Birthdate: params.Birthdate,
		Stack:     params.Stack,
	}

	err = person.Validate()
	if err != nil {
		return nil, err
	}

	return person, nil
}

func BuildPerson(params PersonParams) (person *Person, err error) {
	person = &Person{
		ID:        params.ID,
		Nickname:  params.Nickname,
		Name:      params.Name,
		Birthdate: params.Birthdate,
		Stack:     params.Stack,
	}

	err = person.Validate()
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (e *Person) Validate() (err error) {
	var fieldRules []*validation.FieldRules

	// Attr `Name`
	fieldRules = append(fieldRules, validation.Field(
		&e.Name, // field
		// rules
		validation.Required,
		validation.Length(0, 100),
	))

	// Attr `Nickname`
	fieldRules = append(fieldRules, validation.Field(
		&e.Nickname, // field
		// rules
		validation.Required,
		validation.Length(0, 32),
	))

	// Attr `Birthdate`
	fieldRules = append(fieldRules, validation.Field(
		&e.Birthdate, // field
		// rules
		validation.Required,
		validation.Date("2006-01-02"),
	))

	// Attr `Stack`
	fieldRules = append(fieldRules, validation.Field(
		&e.Stack, // field
		// rules
		validation.Skip.When(e.Stack == nil),
		validation.Each(
			validation.Required,
			validation.Length(0, 32),
		),
	))

	return validation.ValidateStruct(e, fieldRules...)
}
