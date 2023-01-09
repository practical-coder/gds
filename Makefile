.PHONY:all build clean fmt vet scheck
all: build
build: scheck
	go build -o gds
clean:
	rm -f gds
fmt:
	go fmt ./...
test:
	go test -v ./...
vet: fmt
	go vet ./...
scheck: vet
	staticcheck ./...
