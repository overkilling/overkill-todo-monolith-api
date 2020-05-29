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
test:
	@echo "===Unit Tests==="
	go test -cover ./... -short

.PHONY: integration
integration:
	@echo "===Integration Tests==="
	go test -cover ./... -run TestIntegrationRouter

.PHONY: pact
pact:
	@echo "===Pact==="
	$(PACT_DOCKER_COMPOSE) up --build --abort-on-container-exit
	$(PACT_DOCKER_COMPOSE) down --volumes
