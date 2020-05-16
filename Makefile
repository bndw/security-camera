GOARGS=GOOS=linux GOARCH=arm CGO_ENABLED=0

.PHONY: build
build:
	@cd ./uploader \
	&& $(GOARGS) go build -o ../root/usr/local/bin/uploader .
