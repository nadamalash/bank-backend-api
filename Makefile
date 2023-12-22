postgres:
	docker run --name postgres-container -e POSTGRES_USER=nada -e POSTGRES_PASSWORD=koty123 -p 5432:5432 -d postgres:alpine

createdb:
	docker exec -it postgres-container createdb --username=nada --owner=nada bank_db

dropdb:
	docker exec -it postgres-container dropdb bank_db

postgresstop:
	docker stop postgres-container

postgresrm:
	docker rm postgres-container	

migrateup:
	 migrate -path "./db/migrations" -database "postgres://nada:koty123@localhost:5432/bank_db?sslmode=disable" up

migratedown:
	 migrate -path "./db/migrations" -database "postgres://nada:koty123@localhost:5432/bank_db?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb postgresstop postgresrm migrateup migratedown sqlc test server