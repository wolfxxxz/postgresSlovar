create_migration:
    migrate create -ext sql -dir ./schema -seq init
migrate_up:
    migrate -path ./schema -database "postgres://localhost:5435/postgres?sslmode=disable&user=postgres&password=1" up
migrate_down:
    migrate -path ./schema -database "postgres://localhost:5435/postgres?sslmode=disable&user=postgres&password=1" down
docker_run_postgres:
    docker-compose up --build