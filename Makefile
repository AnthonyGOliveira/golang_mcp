# Makefile for Go project

# Default target
all: run test cover

# Target to run the main application
run:
	go run cmd/main.go

# Target to run tests
test:
	go test ./...

# Target to run tests with coverage
cover: cover-test cover-report

cover-test:
	go test ./... -coverprofile=coverage.out

cover-report:
	go tool cover -html=coverage.out

# Target to clean up generated files
clean:
	rm -f coverage.out

# Target to show help (optional)
help:
	@echo "Usage: make <target>"
	@echo "Targets:"
	@echo "  all     Run, test, and generate coverage report (default)"
	@echo "  run     Run the main application"
	@echo "  test    Run tests"
	@echo "  cover   Run tests with coverage and generate report"
	@echo "  cover-test  Runs go test with coverage"
	@echo "  cover-report Generates html coverage report"
	@echo "  clean   Clean up generated files"
	@echo "  help    Show this help message"

.PHONY: all run test cover cover-test cover-report clean help
