.DEFAULT_GOAL := build

.PHONY: fmt
fmt:
	@echo "Formatting project..."
	@go fmt ./...

.PHONY: vet
vet: fmt
	@echo "Vetting the project..."
	@go vet ./...

.PHONY: build // always start with this
build: vet 
	@go build
	@echo "Building the project..." // for commits

.PHONY: clean
clean:
	@rm go-basics
	@echo "Cleaning up..."