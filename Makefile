.PHONY: all build clean check test vet fmt lint

build:
	go build -o start/start.out start/main.go

check: fmt vet lint test

test:
	go test ./...

vet:
	go vet ./...

fmt:
	go fmt ./...

lint:
	golint ./...

clean:
	-rm *.o *.log *.exe *.out
#	-rm start/start