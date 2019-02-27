#!/usr/bin/env bash

set -euxo pipefail

# cd to a project root. {
cd $(dirname $0)
cd ../
# }

# DATABASE is a name of Database.
DATABASE=$1
# USER is a mysql login user.
USER=$2
# PASSWORD is a mysql login password.
PASSWORD=$3
# HOST is a mysql server host.
HOST=$4
# SCHEMA_PATH is a new path for schema.sql.
SCHEMA_PATH=/tmp/schema.sql

cp _sql/mysql/schema.sql "$SCHEMA_PATH"
sed -i "" -e"s/test-xo-db/$DATABASE/g" "$SCHEMA_PATH"
mysql -u"$USER" -p"$PASSWORD" -h "$HOST" < "$SCHEMA_PATH"
rm "$SCHEMA_PATH"
