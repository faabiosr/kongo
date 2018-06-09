.DEFAULT_GOAL := test

# Clean up
clean:
	@rm -fR ./vendor/ ./cover.*
.PHONY: clean

# Creates folders and download dependencies
configure:
	dep ensure -v
.PHONY: configure

# Run tests and generates html coverage file
cover: test
	go tool cover -html=./cover.out -o ./cover.html
.PHONY: cover

# Download dependencies
depend:
	go get -u gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install
	go get -u github.com/golang/dep/...
.PHONY: depend

# Format all go files
fmt:
	gofmt -s -w -l $(shell go list -f {{.Dir}} ./... | grep -v /vendor/)
.PHONY: fmt

# Run linters
lint:
	gometalinter.v2 --vendor ./...
.PHONY: lint

# Run tests
test:
	go test -v -race -coverprofile=./cover.out $(shell go list ./... | grep -v /vendor/)
.PHONY: test
