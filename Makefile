# Dependencies :
# docker
# migrate
# sqlc
# mockgen

all: dev

.PHONY: run-dev
run-dev: dev up_dev migrate
	@./url_shortener

.PHONY: run
run: docker up migrate

.PHONY: stop
stop: down clean

.PHONY: test
test:
	@mockgen -source=./internal/service/url.go -package=urlsvcmock -destination=./internal/service/mocks/url.mock.go
	@go test -cover ./...

.PHONY: dev
dev: clean
	@go mod tidy
	@go build -v -o url_shortener .

.PHONY: docker
docker: clean
	@go mod tidy
	@GOOS=linux GOARCH=arm go build --tags=docker -o url_shortener .
	@docker rmi -f vinchent123/url_shortener:v0.0.1
	@docker build -t vinchent123/url_shortener:v0.0.1 .

.PHONY: clean
clean:
	@rm -f url_shortener

.PHONY: up_dev
up_dev:
	@docker compose -f docker-compose-dev.yml up -d

.PHONY: up
up:
	@docker compose up -d

.PHONY: down
down:
	@docker compose down

.PHONY: migrate
migrate:
	@echo "Wait 3s for Postgres to startup"
	@sleep 3
	@migrate -database "postgres://postgres:postgres@localhost:25432/url_shortener?sslmode=disable" -path ./migrations/ up
	@echo "Run migrations"

.PHONY: migrate-down
migrate-down:
	@migrate -database "postgres://postgres:postgres@localhost:25432/url_shortener?sslmode=disable" -path ./migrations/ down
