package cfg

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pgOnce sync.Once
var DB *pgxpool.Pool

func Pg() *pgxpool.Pool {
	pgOnce.Do(func() {
		url := fmt.Sprintf(
			"postgres://%s:%s@%s/%s?search_path=%s",
			Cfg().Postgres.Username,
			Cfg().Secrets.PostgresPassword,
			Cfg().Postgres.Address,
			Cfg().Postgres.DbName,
			Cfg().Postgres.Schema,
		)
		var err error
		DB, err = pgxpool.New(context.Background(), url)
		if err != nil {
			panic(fmt.Sprintf("Unable to connect to database: %v", err))
		}
	})
	return DB
}
