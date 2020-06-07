.PHONY: build test

build:
	go build -o ./bin/open-dam-worker ./cmd/open-dam-worker/

build-docker:
	docker build -t open-dam-worker -f build/Dockerfile .

test:
	go test -v ./...

test-coverage:
	if [ ! -d coverage ]; then mkdir coverage; fi
	go test -coverpkg ./internal/... -coverprofile coverage/coverage.out ./... && go tool cover -html=coverage/coverage.out

run:
	docker-compose --file ./build/docker-compose.yml --project-directory . up --build
