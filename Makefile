.PHONY: test

test:
	@go test -cover -v ./...

fmt:
	@go fmt ./...