BINARY_NAME=tldr-webdev-tools

prep:
	go mod tidy
	go mod download

build:
	GOARCH=amd64 GOOS=darwin CGO_ENABLED=0 go build -o bin/${BINARY_NAME}-darwin

run: build
	bin/${BINARY_NAME}-darwin