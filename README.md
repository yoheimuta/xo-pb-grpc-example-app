# xo-example-app

[![CircleCI](https://circleci.com/gh/yoheimuta/xo-example-app/tree/master.svg?style=svg)](https://circleci.com/gh/yoheimuta/xo-example-app/tree/master)

## Setup

Set the project root path to GOPATH.

```bash
mkdir xo-example-app
cd xo-example-app
export GOPATH=$(pwd)
go get -d github.com/yoheimuta/xo-example-app # ignore a `no Go files` error.
cd src/github.com/yoheimuta/xo-example-app
```

Install a dependency.

```bash
go get -u github.com/xo/xo
```

Run your mysql.

```bash
docker run --name test-mysql --rm -d -e MYSQL_ROOT_PASSWORD=my-pw -p 3306:3306 mysql:8.0.0
```

Create a database and tables.

```bash
mysql -uroot -p'my-pw' -h 0.0.0.0 < _sql/mysql/schema.sql
```

## Development

Update generated model code.

```bash
make dev/gen/xo
```

## Testing

Run fmt, lint and test.


```
make dev/test/all
```
