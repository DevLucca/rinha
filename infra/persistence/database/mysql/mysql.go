package mysql

import (
	"context"
	"fmt"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/DevLucca/rinha/infra/config"
)

type MySQLClient struct {
	client *sql.DB
}

func NewMySQLClient(ctx context.Context, cfg *config.Config) (dbClient *MySQLClient, err error) {
	dsn := buildDBConnectionURI(cfg)
	fmt.Println(dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLClient{
		client: db,
	}, nil
}

func buildDBConnectionURI(cfg *config.Config) string {
	return fmt.Sprintf("%s:%s@/%s", cfg.Db.User, cfg.Db.Pass, cfg.Db.Name)
}
