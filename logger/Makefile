run:
	go build -o bin/logger . && bin/logger

test-only:
	CGO_ENABLED=0 go test -count=1 ./... -v

test: test-only
