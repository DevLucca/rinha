package inmemory

import (
	"context"
	"sync"

	"github.com/DevLucca/rinha/domain/entity"

	"github.com/google/uuid"
)

type InMemoryPersonRepository struct {
	mu     *sync.RWMutex
	people map[uuid.UUID]*entity.Person
}

func NewInMemoryPersonRepository() *InMemoryPersonRepository {
	return &InMemoryPersonRepository{
		mu:     &sync.RWMutex{},
		people: make(map[uuid.UUID]*entity.Person),
	}
}

func (r *InMemoryPersonRepository) List(ctx context.Context) ([]*entity.Person, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	people := make([]*entity.Person, 0)
	for _, v := range r.people {
		people = append(people, v)
	}
	return people, nil
}

func (r *InMemoryPersonRepository) GetByID(ctx context.Context, uuid uuid.UUID) (*entity.Person, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	person := r.people[uuid]
	return person, nil
}

func (r *InMemoryPersonRepository) Save(ctx context.Context, person *entity.Person) error {
	r.mu.Lock()
	r.people[person.ID] = person
	r.mu.Unlock()
	return nil
}
