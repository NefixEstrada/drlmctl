export GO111MODULE=on
BINARY_NAME=drlmctl

all: deps build
install:
	go install drlmctl.go
build:
	go build -o $(BINARY_NAME) drlmctl.go
test:
	go test -cover ./...
clean:
	go clean
	rm -f $(BINARY_NAME)
deps:
	go build -v ./...
upgrade:
	go get -u
