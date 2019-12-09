default: test

.PHONY: lint
lint:
	cd test && golangci-lint run -v

.PHONY: test
test:
	cd test && go test -v
