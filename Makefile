.PHONY: all build_windows build_linux build_mac clean check test vet fmt lint

build_windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o start/start.exe start/main.go

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o start/start.out start/main.go

build_mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o start/start start/main.go

all: build_windows build_linux build_mac

check: fmt vet lint test

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golint ./...

install:
print:
tar:
dist:
TAGS:


clean:
	$(shell  find . -regex '.*\.out\|.*\.exe\|.*\.i\|.*\.s\|.*\.o' | xargs rm -f)

