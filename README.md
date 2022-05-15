Api Repo for Quetzal project.
Folder structure follows https://github.com/golang-standards/project-layout

# Internal
## server
    creates server with dependencies(db, logger, router)
## transport
    http handler logic. Generally just parses requests and calls service layer

## service
    buisness logic, calls storage + other apis
## storage
    storage layer for db access. Just postgres/redis for now



# Migrations
    To create a new migration file install go-migrate on your machine & then
    `migrate create -ext sql -dir db/migrations -seq your_migration`
    To run migrations in application startup while developing set env variable 
    `LOCAL_MIGRATIONS=true`
    To run migrations outside application startup `migrate -path migrations -database "postgres://localhost:5432/quetzal?sslmode=disable" up 1`
# documentation
    