migrateup:
	docker-compose exec app sh -c 'migrate -database $$DATABASE_URL -path db/migration -verbose up'

migratedown:
	docker-compose exec app sh -c 'migrate -database $$DATABASE_URL -path db/migration -verbose down'

sqlc:
	docker-compose exec app sqlc generate

.PHONY: migrateup migratedown sqlc
