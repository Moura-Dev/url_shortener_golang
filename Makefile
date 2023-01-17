PKGS ?= $(shell go list ./...)
POSTGRESQL_URL ?= postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
.PHONY: all services clean

unit-test:
	go test -covermode=atomic -coverprofile coverage.out -v ${PKGS}  -short -tags=unit

integration-test:
	go test -covermode=atomic -coverprofile coverage.out -v ${PKGS} -tags=integration

down:
	docker-compose down && docker volume prune -f

ensure-dependencies:
	go mod tidy

up:
	docker-compose up -d
