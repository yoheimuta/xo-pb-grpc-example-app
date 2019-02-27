# commands for development {
# ------------------------------------------------------------------
# ------------------------------------------------------------------

## dev/test/all runs fmt, lint and test.
dev/test/all:dev/fmt test/lint dev/test

## dev/fmt enforces a right code.
dev/fmt:
	goimports -w `find . -name vendor -prune -type f -o -name '*.go'`
	gofmt -s -w `find . -name vendor -prune -type f -o -name '*.go'`
	unconvert -apply ./...

## dev/test runs a test.
dev/test:
	go test -v -p 8 -count 1 -timeout 240s -race ./...

## dev/add/gopkg vendors the library specified by a GOPKG variable.
dev/add/gopkg:
	dep ensure -add $(GOPKG)

# ------------------------------------------------------------------
# ------------------------------------------------------------------
# }

# commands for test {
# ------------------------------------------------------------------
# ------------------------------------------------------------------

## test/lint runs lint.
test/lint:
	(! gofmt -s -d `find . -name vendor -prune -type f -o -name '*.go'` | grep '^')
	golint -set_exit_status `go list ./...`
	(! goimports -l `find . -name vendor -prune -type f -o -name '*.go'` | grep 'go')
	go vet ./...
	go vet -shadow ./...
	# TODO: Rewrite xo template.
	# errcheck ./...
	varcheck ./...
	ineffassign .
	unconvert -v ./...

# ------------------------------------------------------------------
# ------------------------------------------------------------------
# }
