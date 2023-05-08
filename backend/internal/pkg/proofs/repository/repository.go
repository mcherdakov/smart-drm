package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/jackc/pgx/pgtype"
	"github.com/jmoiron/sqlx"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) SetProof(ctx context.Context, p entity.Proof, address string) error {
	dbProof := entity.DBProof{
		V:       p.V,
		R:       p.R,
		S:       p.S,
		Hash:    p.Hash,
		Date:    p.Date,
		Value:   p.Value,
		Address: strings.ToLower(address),
	}

	_, err := r.db.NamedExecContext(
		ctx,
		`
            insert into proof(v, r, s, hash, date, value, address)
                values (:v, :r, :s, :hash, :date, :value, :address)
            on conflict(address) do update set
                v = excluded.v,
                r = excluded.r,
                s = excluded.s,
                hash = excluded.hash,
                date = excluded.date,
                value = excluded.value
        `,
		dbProof,
	)

	return err
}

func (r *Repository) ProofByAddress(ctx context.Context, address string) (*entity.DBProof, error) {
	var proof entity.DBProof

	err := r.db.QueryRowxContext(
		ctx,
		`select v, r, s, hash, date, value, address, last_commited_date, last_commited_value from proof
            where address=$1`,
		address,
	).StructScan(&proof)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &proof, nil
}

func (r *Repository) UnsyncedProofsAcquire(ctx context.Context, tx *sqlx.Tx) ([]entity.DBProof, error) {
	rows, err := tx.QueryxContext(
		ctx,
		`
        select v, r, s, hash, date, value, address, last_commited_date, last_commited_value from proof
            where coalesce(last_commited_value, 0) != value
        for update
        `,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	proofs := []entity.DBProof{}

	for rows.Next() {
		var proof entity.DBProof

		if err := rows.StructScan(&proof); err != nil {
			return nil, err
		}

		proofs = append(proofs, proof)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return proofs, nil
}

func (r *Repository) UnsyncedProofsUpdate(ctx context.Context, tx *sqlx.Tx, proofs []entity.DBProof) error {

	addresses := make([]string, 0, len(proofs))
	for _, p := range proofs {
		addresses = append(addresses, p.Address)
	}

	dbAddresses := &pgtype.TextArray{}
	if err := dbAddresses.Set(addresses); err != nil {
		return err
	}

	_, err := tx.ExecContext(
		ctx,
		`
            update proof set
                last_commited_date=date,
                last_commited_value=value
            where address = any($1::text[])
        `,
		dbAddresses,
	)

	return err
}

func (r *Repository) Transaction(ctx context.Context, f func(ctx context.Context, tx *sqlx.Tx) error) error {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	err = f(ctx, tx)
	if err != nil {
		_ = tx.Rollback()

		return err
	}

	return tx.Commit()
}
