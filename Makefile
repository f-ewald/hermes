build:
	env GOOS=darwin GOARCH=amd64 go build -o hermes-macos-amd64 main.go
	env GOOS=darwin GOARCH=arm64 go build -o hermes-macos-arm64 main.go