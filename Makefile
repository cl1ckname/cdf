# Application name
APP_NAME = cdf

# Version
VERSION = 0.1.0

# Go build command
GOOS = $(shell go env GOOS)
GOARCH = $(shell go env GOARCH)
GO_BUILD = go build -ldflags="-X main.version=$(VERSION) -s -w" -o bin/$(APP_NAME)-$(VERSION)-$(GOOS)-$(GOARCH)

# Platforms to build for
PLATFORMS = linux darwin

# Architectures to build for
ARCHITECTURES = amd64 arm64
.PHONY:
	build
	install
	all

all: linter test build

build-all: $(PLATFORMS)

$(PLATFORMS):
	@$(foreach GOARCH, $(ARCHITECTURES), \
		$(foreach GOOS, $@, \
			$(eval BIN = bin/$(APP_NAME)-$(VERSION)-$(GOOS)-$(GOARCH)) \
			echo "Building $(BIN)"; \
			GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD); \
		) \
	)


build:
	@mkdir -p bin
	@BINARY="bin/$(APP_NAME)-$(VERSION)-$(GOOS)-$(GOARCH)"; \
	echo "Building $$BINARY"; \
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD)
	
# Install the binary to $GOBIN
install: build
	@if [ -z "$(GOBIN)" ]; then \
		echo "Error: GOBIN is not set. Please set GOBIN to your desired installation directory."; \
		exit 1; \
	fi
	@if [ ! -d "$(GOBIN)" ]; then \
		echo "Error: GOBIN directory $(GOBIN) does not exist."; \
		exit 1; \
	fi
	@BINARY="bin/$(APP_NAME)-$(VERSION)-$(GOOS)-$(GOARCH)"; \
	cp "$$BINARY" "$(GOBIN)/$(APP_NAME)"; \
	if [ ! -f "$$BINARY" ]; then \
		echo "Error: Binary $$BINARY not found. Please run 'make build' first."; \
		exit 1; \
	fi
	@echo "Installed $$BINARY to $(GOBIN)/$(APP_NAME)"
	
test:
	go test ./...

linter:
	golangci-lint run

clean:
	rm -rf bin/
