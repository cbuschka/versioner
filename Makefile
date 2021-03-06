TOP_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

build:
	@cd ${TOP_DIR} && \
	mkdir -p dist/ && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-extldflags \"-static\"" -o dist/semver ./cmd/semver.go

clean:
	rm -rf ${TOP_DIR}/dist/ ${TOP_DIR}/.cache/

test:
	@cd ${TOP_DIR} && \
	go test -cover -race -coverprofile=coverage.txt -covermode=atomic ./cmd/...
