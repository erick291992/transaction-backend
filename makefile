dockerstart:
	docker-compose -f docker-compose.yml --env-file .env.development up
dockerbash:
	docker-compose -f docker-compose.yml run --service-ports web bash
migrateupall:
	migrate -path  src/db/migration -database "postgresql://admin:secret@localhost:5432/main?sslmode=disable" -verbose up
migratedownall:
	migrate -path  src/db/migration -database "postgresql://admin:secret@localhost:5432/main?sslmode=disable" -verbose down
migrateup:
	migrate -path  src/db/migration -database "postgresql://admin:secret@localhost:5432/main?sslmode=disable" -verbose up 1
migratedown:
	migrate -path  src/db/migration -database "postgresql://admin:secret@localhost:5432/main?sslmode=disable" -verbose down 1
migrateinit:
	migrate create -ext sql -dir src/db/migration -seq init_schema
migratecreate:
	@read -p "Enter migration name. " name; \
		migrate create -ext sql -dir src/db/migration -seq $$name
sqlc:
	sqlc generate