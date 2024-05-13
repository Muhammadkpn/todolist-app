.PHONY: build-http
build-http:
	@go build -v -o bin/http cmd/http/*.go

.PHONY: build-http-non-di
build-http-non-di:
	@go build -v -o bin/httpNonDi cmd/httpNonDi/*.go

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
	@go clean -testcache
	@go test `go list ./...` -cover -coverprofile=coverage.out --race -coverpkg=./internal/usecase/...,./internal/repository/...
	@go tool cover -func=coverage.out

.PHONY: mockgen
mockgen:
	@mockery --all --keeptree --disable-version-string

.PHONY: lint
lint:
	@golangci-lint run -v
