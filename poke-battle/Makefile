DB_URL=postgresql://root:secret@localhost:5432/pokebattle?sslmode=disable

build:
	go build -o bin/pokebattle

run:
	go run main.go

sqlc:
	docker run --rm -v /home/cristian/development/go/poke-battle-app:/src -w /src/poke-battle sqlc/sqlc generate

migrateup:
	migrate -path ../db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path ../db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir ../db/migration -seq $(name)

proto:
	rm -rf pb/*.go && protoc2 --proto_path=../proto --go_out=. --go-grpc_out=. ../proto/*.proto