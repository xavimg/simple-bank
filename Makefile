createdb:
	docker exec -it postgres12 createdb --username=xmg simple_bank

dropdb:
	docker exec -it postgres12 dropdb --username=xmg simple_bank

migrate_up:
	migrate -path db/migration -database "postgresql://xmg:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://xmg:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrate_up migrate_down sqlc test
