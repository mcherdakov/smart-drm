run:
	docker compose up --build

db:
	PGPASSWORD=password psql postgres://127.0.0.1:5434/smartdrm --user node

fixtures:
	PGPASSWORD=password psql postgres://127.0.0.1:5434/smartdrm --user node < backend/fixtures/content.sql

