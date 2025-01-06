#!/bin/sh

set -e

echo "Running database migrations"
/app/migrate -path /app/migration -database "$DATABASE_URL" -verbose up

