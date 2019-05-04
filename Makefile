TARGET = taskar
SOURCE = main.go
DBPATH = ${HOME}/go/data/taskar

all: format build

install: main.go
	@echo "==> Setting up database"
	mkdir -p ${DBPATH}
	@echo "==> Installing Taskar.."
	go install
	@echo "==> Done!"

build: main.go
	@echo "==> Building from source.."
	go build -o ${TARGET} ${SOURCE}

.PHONY: all test format clean

format:
	@echo "==> Formatting code.."
	gofmt -w .

test:
	@echo "==> Running tests.."
	go test ./... -v

clean:
	@echo "==> Cleaning up.."
	go mod tidy
	go clean
	rm -rf ${DBPATH}
