all: install lint lint-fix

install:
	CGO_ENABLED=0 go install -tags 'netgo osusergo' -ldflags='-s -w' ./cmd/dayone2md

lint:
	golangci-lint run
	govulncheck -test .

lint-fix:
	golangci-lint run --fix