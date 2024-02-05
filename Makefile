.PHONY: build-http
build-http:
	@go build -v -o bin/http cmd/http/*.go

.PHONY: run-http
run-http:
	@go run -race -v cmd/http/main.go

.PHONY: docker-build
docker-build:
	@docker-compose build

.PHONY: docker-up
docker-up:
	@docker-compose up

.PHONY: docker-down
docker-down:
	@docker-compose down

.PHONY: test
test:
	@go test -cover -count=1 --race ./...

.PHONY: mockgen
mockgen:
	@go generate ./...