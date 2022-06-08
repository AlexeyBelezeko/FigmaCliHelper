build_mac_arm:
	GOOS=darwin GOARCH=arm64 go build -o bin/figma-cli-arm64-darwin cmd/main.go

build_mac_amd:
	GOOS=darwin GOARCH=amd64 go build -o bin/figma-cli-amd64-darwin cmd/main.go
