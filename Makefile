TARGET = taskar
SOURCE = main.go

.PHONY: all test build execute format clean

all: format build

test:
	@echo "==> Running tests.."
	go test ./... -v

build:
	@echo "==> Building from source.."
	go build -o ${TARGET} ${SOURCE}

format:
	@echo "==> Formatting code.."
	gofmt -w .

clean:
	@echo "==> Cleaning up.."
	go mod tidy
	go clean
	rm data/tasks.db
