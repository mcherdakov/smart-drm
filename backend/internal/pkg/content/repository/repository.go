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
        select id, author, header, content,
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

func (r *Repository) ContentStats(ctx context.Context) ([]entity.ContentClickAggregate, error) {
	rows, err := r.db.QueryxContext(
		ctx,
		`
        select id, author, header,
        (
            select count(*) from click 
                where content_id=id and commited='t'
        ) as clicks_commited,
        (
            select count(*) from click 
                where content_id=id and commited='f'
        ) as clicks_uncommited
        from content
        order by id
        `,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []entity.ContentClickAggregate{}

	for rows.Next() {
		var item entity.ContentClickAggregate

		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repository) AuthorStats(ctx context.Context) ([]entity.AuthorClickAggregate, error) {
	rows, err := r.db.QueryxContext(
		ctx,
		`
        select content.author, click.commited, count(*) from click
            join content on click.content_id = content.id
        group by content.author, click.commited
        `,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []entity.AuthorClickAggregate{}

	for rows.Next() {
		var item entity.AuthorClickAggregate

		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		items = append(items, item)
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
        select id, author, header, content
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
        insert into click(content_id, date, address, commited)
        values(:content_id, :date, :address, :commited)
        on conflict(content_id, date, address) do nothing
        `,
		click,
	)

	return err
}
