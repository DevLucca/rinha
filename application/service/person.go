package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DevLucca/rinha/application/dto"
	"github.com/DevLucca/rinha/domain/entity"
	"github.com/DevLucca/rinha/domain/service"
	"github.com/DevLucca/rinha/infra/persistence/cache"
	"github.com/google/uuid"
)

type PersonService struct {
	cache     cache.Cache
	domainSvc *service.PersonService
}

func NewPersonService(cache cache.Cache, domainSvc *service.PersonService) *PersonService {
	return &PersonService{
		cache:     cache,
		domainSvc: domainSvc,
	}
}

func (svc *PersonService) Create(ctx context.Context, dto dto.PersonRequestDTO) (id uuid.UUID, err error) {
	params := entity.PersonParams{
		Name:      dto.Name,
		Nickname:  dto.Nickname,
		Birthdate: dto.Birthdate,
		Stack:     dto.Stack,
	}

	id, err = svc.domainSvc.Save(ctx, params)
	if err != nil {
		return id, err
	}

	params.ID = id
	jsonBytes, _ := json.Marshal(params)

	svc.cache.SetItem(fmt.Sprintf("person:%s", id.String()), jsonBytes)
	svc.cache.Increase("person-count")
	return id, nil
}

func (svc *PersonService) Retrieve(ctx context.Context, id uuid.UUID) (person *entity.Person, err error) {
	jsonBytes, err := svc.cache.GetItem(fmt.Sprintf("person:%s", id.String()))
	if err != nil {
		if err == cache.ErrCacheMiss {
			person, err = svc.domainSvc.Retrieve(ctx, id)
			if err != nil {
				return nil, err
			}
			svc.cache.SetItem(fmt.Sprintf("person:%s", id.String()), jsonBytes)
			return person, err
		}
		return nil, err
	}

	json.Unmarshal(jsonBytes, &person)
	return person, nil
}

func (svc *PersonService) Count(ctx context.Context) int {
	count, _ := svc.cache.GetInt("person-count")
	return int(count)
}
