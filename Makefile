BUILD_DIR := bin
TOOLS_DIR := tools

#Variables de entorno
HOST := localhost
PORT := 1323
APIKEY := cur_live_mhcdXGJOTpPgfyrnE5WWXxsGAysjzHpzvQJT5HOg
HP_POSTGRES_HOST := localhost
HP_POSTGRES_PORT := 5432
HP_POSTGRES_USER := pguser
HP_POSTGRES_PASSWORD := oDIx9eGKzlwqYGeE
HP_POSTGRES_DB := test-boletia-db
INTERVAL_MINUTES := 1
TIME_OUT_SECONDS := 5

default: all

all: clean lint test build run

.PHONY: $(BUILD_DIR)/server
bin/server: cmd/server/*.go
	CGO_ENABLED=0 go build -mod vendor -ldflags="-s -w" -o ./bin/server ./cmd/server/

.PHONY: build
build: bin/server

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(TOOLS_DIR)
	@go mod vendor
	@go mod tidy

.PHONY: run
run: build 
	APIKEY=$(APIKEY) TIME_OUT_SECONDS=$(TIME_OUT_SECONDS) INTERVAL_MINUTES=$(INTERVAL_MINUTES) HOST=$(HOST) PORT=$(PORT) HP_POSTGRES_HOST=$(HP_POSTGRES_HOST) HP_POSTGRES_PORT=$(HP_POSTGRES_PORT) HP_POSTGRES_USER=$(HP_POSTGRES_USER) HP_POSTGRES_PASSWORD=$(HP_POSTGRES_PASSWORD) HP_POSTGRES_DB=$(HP_POSTGRES_DB)  bin/server

tools/golangci-lint/golangci-lint:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b tools/golangci-lint latest

.PHONY: lint
lint: $(TOOLS_DIR)/golangci-lint/golangci-lint
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run ./...

.PHONY: test
test:
	go test -mod vendor -race -cover -coverprofile=coverage.txt -covermode=atomic ./...
