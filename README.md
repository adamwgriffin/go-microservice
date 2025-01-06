# Go Microservice

A boilerplate for bootstrapping a new microservice with Go and Gin with sqlc.

## Setup

- [Install Docker](https://www.docker.com/products/docker-desktop)
- Run migrations: `make migrateup`

## Working in Dev

### Run the app
`docker compose up`

### Generate code for SQL queries
`make sqlc`
