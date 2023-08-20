build:
	@GO111MODULE=on go build -o bin/indexer cmd/main.go

clean:
	@rm -rf bin/*

lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1 \
	&& golangci-lint run

dev:
	@RUN_TESTS=false ./scripts/local_dev.sh

test:
	@RUN_TESTS=true ./scripts/local_dev.sh

gen_bindings:
	@./scripts/gen_bindings.sh

.PHONY: build \
				clean \
				lint \
				dev \
				test \
				gen_bindings
