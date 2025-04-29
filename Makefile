BIN_DIR := bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint


default: pregit build

.PHONY: lint
lint: $(GOLANGCI_LINT)
	@$(GOLANGCI_LINT) run

update-golangci:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

format:
	@$(GOLANGCI_LINT) run --fix

release-notes:
	git-cliff > CHANGELOG.md

pregit: format

cargo-brew:
	brew install rust

cliff-install:
	cargo install git-cliff

build:
	go build -o ./miso-server ./apps/server
	go build -o ./miso-runner ./apps/runner
	go build -o ./miso ./apps/cli

clean:
	rm ./miso-server
	rm ./miso-runner
	rm ./miso