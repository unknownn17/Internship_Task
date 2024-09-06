run:
	go run cmd/main.go
table:
	migrate create -dir migrations -ext sql db
up:
	migrate -path migrations -database "postgres://postgres:2005@postgres:5432/internship?sslmode=disable" up
down:
	migrate -path migrations -database "postgres://postgres:2005@postgres:5432/internship?sslmode=disable" down
force:
	migrate -path migrations -database "postgres://postgres:2005@postgres:5432/internship?sslmode=disable" force
swagger:
	swag init -g ./internal/api/router/router.go -o internal/docs