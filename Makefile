migrate create -ext sql -dir ./schema -seq init

migrate -path ./schema -database "postgres://localhost:5435/postgres?sslmode=disable&user=postgres&password=1" up

migrate -path ./schema -database "postgres://localhost:5435/postgres?sslmode=disable&user=postgres&password=1" down

docker-compose up --build

go build -o SlovarPostgres