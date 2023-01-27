POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=developer
POSTGRES_PASSWORD=2002
POSTGRES_DATABASE=coursebot

-include .env
  
DB_URL="postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable"

print:
	echo "$(DB_URL)"
	
migrate_file:
	migrate create -ext sql -dir migrations -seq create_courses_table

run:
	go run cmd/main.go

migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1

.PHONY: start migrateup migratedown
