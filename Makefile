protoc:
	protoc model.proto  --go_out=.
.PHONY: protoc

test:
	go test -v -coverprofile=cover.out ./...
.PHONY: test
