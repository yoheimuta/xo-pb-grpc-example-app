# xo-pb-grpc-example-app

[![CircleCI](https://circleci.com/gh/yoheimuta/xo-pb-grpc-example-app/tree/master.svg?style=svg)](https://circleci.com/gh/yoheimuta/xo-pb-grpc-example-app/tree/master)

## Setup

Set the project root path to GOPATH.

```bash
mkdir xo-pb-grpc-example-app
cd xo-pb-grpc-example-app
export GOPATH=$(pwd)
go get -d github.com/yoheimuta/xo-pb-grpc-example-app # ignore a `no Go files` error.
cd src/github.com/yoheimuta/xo-pb-grpc-example-app
```

Install a dependency.

```bash
go get -u github.com/xo/xo
```

- And get a protoc binary manually.

Run your mysql.

```bash
docker run --name test-mysql --rm -d -e MYSQL_ROOT_PASSWORD=my-pw -p 3306:3306 mysql:8.0.0
```

Create a database and tables.

```bash
mysql -uroot -p'my-pw' -h 0.0.0.0 < _sql/mysql/schema.sql
```

## API Server

Run an API Server.

```bash
go run github.com/yoheimuta/xo-pb-grpc-example-app/cmds/server/api
```

You can ping your API Server after running the server.

```bash
go run github.com/yoheimuta/xo-pb-grpc-example-app/cmds/devel/ping/api
```

## Development

Update generated model code.

```bash
make dev/gen/xo
```

Update generated proto files.

```bash
make dev/gen/proto
```

## Testing

Run fmt, lint and test.


```
make dev/test/all
```
