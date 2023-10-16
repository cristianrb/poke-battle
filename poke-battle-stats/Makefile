DB_URL=postgresql://root:secret@localhost:5432/pokebattle?sslmode=disable

run:
	go run main.go

sqlc:
	docker run --rm -v /home/cristian/development/go/poke-battle-app:/src -w /src/poke-battle-stats sqlc/sqlc generate

proto:
	rm -rf pb/*.go && protoc2 --proto_path=../proto --go_out=. --go-grpc_out=. ../proto/*.proto