.SILENT:
.EXPORT_ALL_VARIABLES:
.PHONY: all test run

all: test run

test:
	go test -failfast -coverprofile=coverage.dev -count=1 confinio/...
	go tool cover -func coverage.dev

run:
	go run -race .