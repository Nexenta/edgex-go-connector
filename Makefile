CLIENT_API := v1beta1

.PHONY: all
all: help

build:
    GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -ldflags '$(LDFLAGS)' -o ./bin/s3xclient-$(CLIENT_API) cmd/$(CLIENT_API)/main.go
test:
    go test -timeout 30s github.com/Nexenta/edgex-go-connector -v
suite:
    go test -timeout 30s github.com/Nexenta/edgex-go-connector/tests/e2e/kv -run TestEnd2EndKVTestSuite -v
