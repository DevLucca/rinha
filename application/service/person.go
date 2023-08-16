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
	if exists := svc.cache.Exists(fmt.Sprintf("nickname:%s", dto.Nickname)); exists {
		return id, fmt.Errorf("person already exists")
	}

	params := entity.PersonParams{
		ID:        uuid.New(),
		Name:      dto.Name,
		Nickname:  dto.Nickname,
		Birthdate: dto.Birthdate,
		Stack:     dto.Stack,
	}

	go func() {
		err = svc.domainSvc.Save(ctx, params)
		if err != nil {
			fmt.Println(err)
			return
		}
		svc.cache.Increase("person-count")
		svc.cache.SetInt(fmt.Sprintf("nickname:%s", params.Nickname), 1)
	}()

	go func() {
		jsonBytes, _ := json.Marshal(params)
		svc.cache.SetItem(fmt.Sprintf("person:%s", id.String()), jsonBytes)
	}()

	return params.ID, nil
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

func (svc *PersonService) Search(ctx context.Context, query string) (people []*entity.Person, err error) {
	peopleIDs, err := svc.domainSvc.Search(ctx, query)
	if err != nil {
		return people, err
	}

	for _, id := range peopleIDs {
		person, err := svc.Retrieve(ctx, id)
		if err != nil {
			return people, err
		}
		people = append(people, person)
	}

	return people, nil
}

func (svc *PersonService) Count(ctx context.Context) int {
	count, _ := svc.cache.GetInt("person-count")
	return int(count)
}
