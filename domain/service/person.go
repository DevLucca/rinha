package service

import (
	"context"

	"github.com/DevLucca/rinha/domain/entity"
	"github.com/DevLucca/rinha/domain/repository"

	"github.com/google/uuid"
)

type PersonService struct {
	repository repository.Person
}

func NewPersonService(repo repository.Person) *PersonService {
	return &PersonService{
		repository: repo,
	}
}

func (svc *PersonService) Save(ctx context.Context, params entity.PersonParams) (err error) {
	person, err := entity.NewPerson(params)
	if err != nil {
		return err
	}

	err = svc.repository.Save(ctx, person)
	if err != nil {
		return err
	}

	return nil
}

func (svc *PersonService) Search(ctx context.Context, query string) (people []uuid.UUID, err error) {
	people, err = svc.repository.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	return people, nil
}

func (svc *PersonService) Retrieve(ctx context.Context, id uuid.UUID) (person *entity.Person, err error) {
	person, err = svc.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return person, nil
}
