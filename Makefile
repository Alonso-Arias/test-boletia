BUILD_DIR := bin
TOOLS_DIR := tools

#Variables de entorno
HOST := localhost
PORT := 1323
API_KEY := cur_live_mhcdXGJOTpPgfyrnE5WWXxsGAysjzHpzvQJT5HOg
API_URL := https://api.currencyapi.com/v3/latest 
#API_URL := https://httpstat.us/200
HP_POSTGRES_HOST := localhost
HP_POSTGRES_PORT := 5433
HP_POSTGRES_USER := postgres
HP_POSTGRES_PASSWORD := 123456
HP_POSTGRES_DB := test_boletia_db
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
	API_URL=$(API_URL) API_KEY=$(API_KEY) TIME_OUT_SECONDS=$(TIME_OUT_SECONDS) INTERVAL_MINUTES=$(INTERVAL_MINUTES) HOST=$(HOST) PORT=$(PORT) HP_POSTGRES_HOST=$(HP_POSTGRES_HOST) HP_POSTGRES_PORT=$(HP_POSTGRES_PORT) HP_POSTGRES_USER=$(HP_POSTGRES_USER) HP_POSTGRES_PASSWORD=$(HP_POSTGRES_PASSWORD) HP_POSTGRES_DB=$(HP_POSTGRES_DB)  bin/server

tools/golangci-lint/golangci-lint:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b tools/golangci-lint latest

.PHONY: lint
lint: $(TOOLS_DIR)/golangci-lint/golangci-lint
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run ./...

.PHONY: test
test:
	go test -mod vendor -race -cover -coverprofile=coverage.txt -covermode=atomic ./... -ldflags="-X main.API_URL=$(API_URL) -X main.API_KEY=$(API_KEY) -X main.TIME_OUT_SECONDS=$(TIME_OUT_SECONDS) -X main.INTERVAL_MINUTES=$(INTERVAL_MINUTES) -X main.HOST=$(HOST) -X main.PORT=$(PORT) -X main.HP_POSTGRES_HOST=$(HP_POSTGRES_HOST) -X main.HP_POSTGRES_PORT=$(HP_POSTGRES_PORT) -X main.HP_POSTGRES_USER=$(HP_POSTGRES_USER) -X main.HP_POSTGRES_PASSWORD=$(HP_POSTGRES_PASSWORD) -X main.HP_POSTGRES_DB=$(HP_POSTGRES_DB)"

