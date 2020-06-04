BIN_DIR := $(GOPATH)/bin
GOTEST := $(BIN_DIR)/gotest
PACT_DOCKER_COMPOSE := docker-compose -f pact/docker-compose.yml

.PHONY: all
all: test integration pact

.PHONY: run
run:
	go run cmd/todoapi/main.go

.PHONY: build
build:
	go build -o todoapi cmd/todoapi/main.go

.PHONY: test
test: ${GOTEST}
	@echo "===Unit Tests==="
	gotest -cover ./... -short

.PHONY: integration
integration: ${GOTEST}
	@echo "===Integration Tests==="
	gotest -cover ./... -run TestGetAllTodos

.PHONY: pact
pact:
	@echo "===Pact==="
	$(PACT_DOCKER_COMPOSE) up --build --abort-on-container-exit
	$(PACT_DOCKER_COMPOSE) down --volumes

$(GOTEST):
	go get -u github.com/rakyll/gotest
