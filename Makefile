.PHONY: generate test lint build fmt vet clean

# Generate API clients from OpenAPI schemas
generate:
	go run ./cmd/codegen -schema schemas/forum.json -output forum -package forum
	go run ./cmd/codegen -schema schemas/market.json -output market -package market

# Run all tests
test:
	go test -race -count=1 ./...

# Run linters
lint: vet fmt

# Run go vet
vet:
	go vet ./...

# Check formatting
fmt:
	@diff=$$(gofmt -d .); \
	if [ -n "$$diff" ]; then \
		echo "$$diff"; \
		echo "Run: gofmt -w ."; \
		exit 1; \
	fi

# Format code
fmt-fix:
	gofmt -w .

# Build all packages
build:
	go build ./...

# Clean build artifacts
clean:
	go clean ./...
