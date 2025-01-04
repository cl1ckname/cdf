.PHONY:
	build
	install

build:
	mkdir -p bin
	go build -o bin/cdf main.go 
	
install: build
	cp bin/cdf $(GOBIN)/cdf
