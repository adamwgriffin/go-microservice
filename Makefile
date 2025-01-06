migrateup:
	docker-compose exec app sh -c './migrate -database $$DATABASE_URL -path ./migration -verbose up'

migratedown:
	docker-compose exec app sh -c './migrate -database $$DATABASE_URL -path ./migration -verbose down'

sqlc:
	sqlc generate

.PHONY: migrateup migratedown sqlc
