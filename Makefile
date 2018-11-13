
.PHONY: build
build:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ./bin/cctv_upload ./cmd/cctv_upload
