.PHONY: all
all: test integration pact

.PHONY: run
run:
	go run main.go

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
	docker-compose -f docker-compose.pact.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.pact.yml down --volumes