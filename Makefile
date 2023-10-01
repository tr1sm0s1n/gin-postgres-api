BIN_LOC=bin/gin-postgres-api

build:
	docker compose build

up:
	docker compose up -d

down:
	docker compose down

enter:
	docker exec -it gin-postgres psql -d gin-postgres -U demystif -W

compile:
	go build -o ${BIN_LOC} main.go

start:
	./${BIN_LOC}

run:
	go run .

air:
	air
