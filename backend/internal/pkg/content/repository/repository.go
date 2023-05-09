package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/content/entity"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Content(ctx context.Context, date string, address string) ([]entity.Content, error) {
	rows, err := r.db.QueryxContext(
		ctx,
		`
        select id, author, header, content, last_commited_clicks,
            (
                exists (
                    select 1 from click where
                        content_id=id and
                        date=$1 and
                        address=$2
                )
            ) as click_exists
        from content
        order by id
        `,
		date, address,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []entity.Content{}

	for rows.Next() {
		var content entity.Content

		if err := rows.StructScan(&content); err != nil {
			return nil, err
		}

		items = append(items, content)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repository) ContentByID(ctx context.Context, id int64) (*entity.Content, error) {
	var content entity.Content

	err := r.db.QueryRowxContext(
		ctx,
		`
        select id, author, header, content, last_commited_clicks
            from content
        where id=$1
        `,
		id,
	).StructScan(&content)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (r *Repository) CreateClick(ctx context.Context, click entity.Click) error {
	_, err := r.db.NamedExecContext(
		ctx,
		`
        insert into click(content_id, date, address)
            values(:content_id, :date, :address)
        on conflict(content_id, date, address) do nothing
        `,
		click,
	)

	return err
}
