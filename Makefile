.PHONY: help

coverage:
	go test -covermode=count -coverprofile=coverage.out
	go tool cover -html=coverage.out

test:
	go test

