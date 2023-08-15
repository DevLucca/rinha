package service

import (
	"context"
	"fmt"

	"github.com/DevLucca/rinha/domain/entity"
	"github.com/DevLucca/rinha/domain/repository"

	"github.com/google/uuid"
)

type PersonService struct {
	repository repository.Person
}

func NewPersonService(repo repository.Person) *PersonService {
	fmt.Println("oi")
	return &PersonService{
		repository: repo,
	}
}

func (svc *PersonService) Save(ctx context.Context, params entity.PersonParams) (id uuid.UUID, err error) {
	person, err := entity.NewPerson(params)
	if err != nil {
		return id, err
	}

	err = svc.repository.Save(ctx, person)
	if err != nil {
		return id, err
	}

	return person.ID, nil
}

func (svc *PersonService) List(ctx context.Context) {
	svc.repository.List(ctx)
}

func (svc *PersonService) Retrieve(ctx context.Context, id uuid.UUID) (person *entity.Person, err error) {
	person, err = svc.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return person, nil
}
