BIN_DIR := $(GOPATH)/bin
GOTEST := $(BIN_DIR)/gotest
PACT_DOCKER_COMPOSE := docker-compose -f pact/docker-compose.yml

.PHONY: ci
all: test pact

.PHONY: run
run:
	go run cmd/todoapi/*.go

.PHONY: build
build:
	go build -o todoapi cmd/todoapi/*.go

.PHONY: only_unit
only_unit: ${GOTEST}
	@echo "===Only Unit Tests==="
	gotest -cover ./... -short

.PHONY: test
test: ${GOTEST}
	@echo "===All Tests==="
	gotest -cover ./...

.PHONY: pact
pact:
	@echo "===Pact==="
	$(PACT_DOCKER_COMPOSE) up --build --abort-on-container-exit
	$(PACT_DOCKER_COMPOSE) down --volumes

$(GOTEST):
	go get -u github.com/rakyll/gotest
