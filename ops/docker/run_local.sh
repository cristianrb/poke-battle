#!/bin/zsh

export RABBITMQ_URL=amqp://guest:guest@localhost:5672/
export DATABASE_URL=postgresql://root:secret@localhost:5432/pokebattle?sslmode=disable
export POKE_BATTLE_ADDR=:8080
export POKE_BATTLE_STATS_ADDR=localhost:50050

docker-compose up -d

cd ../../poke-battle-stats && go run main.go &
cd ../poke-battle && go run main.go &
cd ../poke-long-battle && go run main.go
