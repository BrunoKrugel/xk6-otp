.PHONY: lint
lint:
	golangci-lint run

.PHONY: build-k6
build-k6:
	xk6 build --with github.com/BrunoKrugel/xk6-otp@latest

.PHONY: format
format:
	goimports -w .
	go fmt ./...
	fieldalignment -fix ./...
