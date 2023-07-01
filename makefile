PROJ=final-project

export APP_NAME=$(PROJ)
export APP_ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export CONFIG_LOCAL=config-local
export CONFIG_DOCKER=config-docker
export DEBUG=true

export APP_CONFIG=$(CONFIG_LOCAL)

build:
	go build -race -o ./$(APP_NAME) ./cmd/main.go

run:
	go run -race ./cmd/main.go

test:
	go test -v -race ./... -count=1

format:
	find . -name '*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

fmt: format

clean:
	rm -f ./$(APP_NAME)

docker:
	docker-compose up --build
