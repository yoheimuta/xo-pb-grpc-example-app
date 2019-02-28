# commands for development {
# ------------------------------------------------------------------
# ------------------------------------------------------------------

## dev/test/all runs fmt, lint and test.
dev/test/all:dev/fmt test/lint test/go

## dev/fmt enforces a right code.
dev/fmt:
	goimports -w `find . -name vendor -prune -type f -o -name '*.go'`
	gofmt -s -w `find . -name vendor -prune -type f -o -name '*.go'`
	unconvert -apply ./...

## dev/install/dep installs dependencies.
dev/install/dep:
	./_script/install_dep.sh

## dev/add/gopkg vendors the library specified by a GOPKG variable.
dev/add/gopkg:
	dep ensure -add $(GOPKG)

## dev/gen/xo generates xo models.
dev/gen/xo:
	xo 'mysql://root:my-pw@0.0.0.0/test-xo-db' \
	    -o infra/expmysql/expmodels \
	    --template-path _xo/templates
	sed -i "" -e"s/test-xo-db\.//g" infra/expmysql/expmodels/*

dev/gen/proto:
	./_proto/BUILD

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
	errcheck ./...
	varcheck ./...
	ineffassign .
	unconvert -v ./...

## test/go runs Go testing.
test/go:
	go test -v -p 8 -count 1 -timeout 240s -race ./...


# ------------------------------------------------------------------
# ------------------------------------------------------------------
# }

