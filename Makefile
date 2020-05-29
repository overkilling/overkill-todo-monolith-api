PACT_DOCKER_COMPOSE := docker-compose -f docker-compose.pact.yml

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
	go test -cover ./...

.PHONY: integration
integration:
	@echo "===Integration Tests==="
	docker-compose up -d db
	go test -cover ./... -tags integration
	docker-compose down

.PHONY: pact
pact:
	@echo "===Pact==="
	$(PACT_DOCKER_COMPOSE) up --build --abort-on-container-exit
	$(PACT_DOCKER_COMPOSE) down --volumes
