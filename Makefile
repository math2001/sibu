.PHONY: help coverage-report
.SILENT:

coverage:
	go test -cover

coverage-report:
	go test -covermode=count -coverprofile=coverage.out
	go tool cover -html=coverage.out

test:
	go test

