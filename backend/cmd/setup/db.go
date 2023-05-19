package setup

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DB() (*sqlx.DB, func() error, error) {
	conn := fmt.Sprintf(
		`
        host=%s
        port=%s
        user=%s
        password=%s
        dbname=%s
        sslmode=disable
        `,
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB"),
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, nil, err
	}

	sqlxDB := sqlx.NewDb(db, "postgres")

	return sqlxDB, sqlxDB.Close, nil
}
