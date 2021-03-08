PROJECT_NAME=$(shell basename "$(PWD)")

DEBUG = false
ifeq ($(DEBUG), true)
DEBUG_FLAGS = -race -v -n
endif

build:
	go build $(DEBUG_FLAGS) -o ./bin/main ./cmd/main.go

run:
	go run $(DEBUG_FLAGS) ./cmd/main.go
