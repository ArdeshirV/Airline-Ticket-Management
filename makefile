PROJ=final-project

export APP_NAME=$(PROJ)
export APP_ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true

build:
	go build -race -o ./$(APP_NAME) ./cmd/main.go

run:
	go run -race ./cmd/main.go

test:
	go test -v -race ./... -count=1

fmt:
	find . -name '*.go' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

clean:
	rm -rf ./$(APP_NAME)

# Docker

