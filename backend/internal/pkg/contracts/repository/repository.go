package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/contracts/entity"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) ContractInsert(ctx context.Context, contract entity.Contract) error {
	_, err := r.db.NamedExecContext(
		ctx,
		`
            insert into contract(name, address)
                values (:name, :address)
        `,
		contract,
	)

	return err
}

func (r *Repository) ContractByName(ctx context.Context, name string) (*entity.Contract, error) {
	var contract entity.Contract

	err := r.db.QueryRowxContext(
		ctx,
		`select name, address from contract
            where name=$1`,
		name,
	).StructScan(&contract)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &contract, nil
}
