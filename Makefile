NAME = deck
TAG  = latest

.DEFAULT_GOAL := test-all

CLI_DOCS_PATH=docs/cli-docs/
.PHONY: test-all
test-all: lint start-test-env docker-test

.PHONY: build-image
build-image:
	docker build -t ${NAME}:${TAG} -f tests/Dockerfile.tests .

.PHONY: start-test-env
start-test-env: build-image
	./tests/init.sh

.PHONY: docker-test
docker-test:
	docker exec deck_deck_1 make test

.PHONY: docker-coverage
docker-coverage:
	docker exec deck_deck_1 make coverage

.PHONY: test
test:
	go test -race ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: build
build:
	CGO_ENABLED=0 go build -o deck main.go

.PHONY: verify-codegen
verify-codegen:
	./scripts/verify-codegen.sh
	./scripts/verify-deepcopy-gen.sh

.PHONY: update-codegen
update-codegen:
	./scripts/update-deepcopy-gen.sh
	go generate ./...

.PHONY: coverage
coverage:
	go test -race -v -count=1 -coverprofile=coverage.out.tmp ./...
	# ignoring generated code for coverage
	grep -E -v 'generated.deepcopy.go' coverage.out.tmp > coverage.out
	rm -f coverage.out.tmp

generate-cli-docs:
	mkdir -p $(CLI_DOCS_PATH)
	go run docs/generate-docs.go -output-path $(CLI_DOCS_PATH)
