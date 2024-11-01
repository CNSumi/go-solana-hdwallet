.PHONY: dev fmt

fmt:
	go mod tidy
	go fmt ./...
	goimports -w -local github.com/cnsumi .

test:
	go test -v ./... -count 1
