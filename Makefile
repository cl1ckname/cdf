build:
	mkdir -p bin
	go build -o bin/cdf cmd/main.go 
install: build
	cp bin/cdf $(GOBIN)/cdf
