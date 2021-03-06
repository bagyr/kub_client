export GO111MODULE=on

.PHONY: all
all: test build

.PHONY: build
build:
	go build -o bin/kub-client github.com/bagyr/kub_client/cmd/server

.PHONY: run
run:
	go run github.com/bagyr/kub_client/cmd/server

.PHONY: run-race
run-race:
	go run -race cmd/kub-client/main.go

.PHONY: test
test:
	go test ./...
#
#.PHONY: docker-run
#docker-run:
#	docker build --build-arg PORT=8080 -t bagyr/banner-api .
#	docker run --rm -p 8080:8080 bagyr/banner-api