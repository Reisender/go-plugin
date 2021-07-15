.PHONY: test clean

test:
	@go generate
	@go test ./...

clean:
	@rm ./example/some-plugin.so || true
