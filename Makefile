e2e:
	go test -v ./... -tags e2e
.PHONY: e2e

unit:
	go test -v ./...
.PHONY: unit
