.PHONY:all build clean fmt vet scheck
all: build
build: scheck
	go build -o gds
clean:
	rm -f gds
fmt:
	go fmt ./...
vet: fmt
	go vet ./...
scheck: vet
	staticcheck ./...
