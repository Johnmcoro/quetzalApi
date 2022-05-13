Api Repo for Quetzal project.
Folder structure follows https://github.com/golang-standards/project-layout

# Internal
## server
    creates server with dependencies(db, logger, router)
## storage
    storage layer for db access. Just postgres/redis for now
## transport
    http handler logic. Generally just calls service
## service
    buisness logic, mainly calls storage + other apis

# Migrations
    To create a new migration file install go-migrate on your machine & then
    `migrate create -ext sql -dir db/migrations -seq your_migration`