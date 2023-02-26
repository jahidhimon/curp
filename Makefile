GOFILES := $(shell find . -name '*.go' -print)
MAIN := .

.PHONY: build run clean

build: $(GOFILES)
	mkdir -p bin
	go build -o bin/curp .

run: 
	go run $(MAIN)

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/curp-linux-arm $(MAIN)
	GOOS=linux GOARCH=arm64 go build -o bin/curp-linux-arm64 $(MAIN)
	GOOS=freebsd GOARCH=386 go build -o bin/curp-freebsd-386 $(MAIN)
	GOOS=windows GOARCH=amd64 go build -o bin/curp-windows-amd64 $(MAIN)

all: hello build

clean:
	rm -rf ./bin
