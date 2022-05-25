BINARY_NAME=bin/gkeeper

all: build

build:
	go build -o ${BINARY_NAME} cmd/goalkeeper/*.go

run:
	go run cmd/goalkeeper/*.go

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test -v