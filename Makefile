DB_URL=postgresql://root:P@ssword123@localhost:5432/social_network?sslmode=disable

postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=P@ssword123 -d postgres:14.2-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root social_network

dropdb:
	docker exec -it postgres dropdb social_network

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server