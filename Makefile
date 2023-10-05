.PHONY: test

test:
	@go test -v ./...

.PHONY: build
build:
	@go build -o lucky .

.PHONY: generate-key
generate-key: build
	@./lucky draw --disable-check

.PHONY: generate-key-full
generate-key-full:
	@./lucky draw