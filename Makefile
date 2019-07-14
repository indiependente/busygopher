all: deps-init deps lint build test run

deps-init:
	rm -f go.mod go.sum
	GO111MODULE=on go mod init
	GO111MODULE=on go mod tidy

deps:
	GO111MODULE=on go mod download

build:
	GO111MODULE=on go build .

lint:
	command -v golangci-lint || (cd /usr/local ; wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest)
    GO111MODULE=on golangci-lint run --disable-all \
    --deadline=10m \
    --skip-files \.*_mock\.*\.go \
    -E errcheck \
    -E govet \
    -E unused \
    -E gocyclo \
    -E golint \
    -E varcheck \
    -E structcheck \
    -E maligned \
    -E ineffassign \
    -E interfacer \
    -E unconvert \
    -E goconst \
    -E gosimple \
    -E staticcheck \
    -E gosec

test:
	GO111MODULE=on go test -race -cover ./...

run:
	GO111MODULE=on go run .