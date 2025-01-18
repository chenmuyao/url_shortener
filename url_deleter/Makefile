all: dev

.PHONY: run-dev
run-dev: dev
	@./url_deleter

.PHONY: dev
dev: clean
	@go mod tidy
	@go build -v -o url_deleter .

.PHONY: docker
docker: clean
	@go mod tidy
	@GOOS=linux GOARCH=arm go build --tags=docker -o url_deleter .
	@docker rmi -f vinchent123/url_deleter:v0.0.1
	@docker build -t vinchent123/url_deleter:v0.0.1 .

.PHONY: clean
clean:
	@rm -f url_deleter
