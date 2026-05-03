lint: check-compatible-versions
	golangci-lint run

fmt:
	go fix ./...
	golangci-lint fmt

fmtlint: fmt lint

test:
	go test -v ./...
