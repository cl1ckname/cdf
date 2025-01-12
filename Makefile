VERSION=0.1.0

.PHONY:
	build
	install

build:
	mkdir -p bin
	go build -ldflags="-X main.version=$(VERSION)" -o bin/cdf main.go 
	
install: build
	cp bin/cdf $(GOBIN)/cdf

test:
	go test ./...

linter:
	golangci-lint run
