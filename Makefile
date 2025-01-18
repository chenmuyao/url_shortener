# Dependencies :
# docker

all: dev

.PHONY: run
run: dev
	@./url_shortener

.PHONY: test
test:
	@go test ./...

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

clean:
	@rm -f url_shortener
