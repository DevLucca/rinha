package repository

import (
	"context"

	"github.com/DevLucca/rinha/domain/entity"

	"github.com/google/uuid"
)

type Person interface {
	Save(context.Context, *entity.Person) error
	GetByID(context.Context, uuid.UUID) (*entity.Person, error)
	Search(context.Context, string) ([]uuid.UUID, error)
}
