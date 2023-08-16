package mysql

import (
	"context"
	"fmt"
	"strings"

	"github.com/DevLucca/rinha/domain/entity"
	"github.com/google/uuid"
)

const (
	peopleTableName = "people"
)

type MySQLPeopleRepository struct {
	client *MySQLClient
}

func NewMySQLPeopleRepository(client *MySQLClient) *MySQLPeopleRepository {
	return &MySQLPeopleRepository{
		client: client,
	}
}

func (r *MySQLPeopleRepository) Save(ctx context.Context, person *entity.Person) (err error) {

	tx, err := r.client.client.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tx.Rollback()

	statement := fmt.Sprintf(`INSERT INTO %s VALUES (?, ?, ?, ?)`, peopleTableName)

	_, err = tx.Exec(statement, person.ID, person.Name, person.Nickname, person.Birthdate)
	if err != nil {
		return err
	}

	statement = `INSERT INTO person_stack VALUES (?, ?)`
	for _, stack := range person.Stack {
		_, err = tx.Exec(statement, person.ID, stack)
		if err != nil {
			return err
		}
	}

	tx.Commit()

	return nil
}

func (r *MySQLPeopleRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Person, error) {
	return nil, nil
}

func (r *MySQLPeopleRepository) Search(ctx context.Context, query string) (people []uuid.UUID, err error) {
	tx, err := r.client.client.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tx.Rollback()

	statement := fmt.Sprintf(`select distinct pp.id from people pp inner join person_stack p on (pp.id = p.id) where concat(pp.name, pp.nickname, p.stack) like '|%s|'`, query)

	statement = strings.ReplaceAll(statement, "|", "%")

	rows, err := tx.Query(statement)
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return
		}

		uuid, _ := uuid.Parse(id)
		people = append(people, uuid)
	}
	return
}
