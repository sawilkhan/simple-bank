#! /bin/sh

set -e

echo "run db migration"
source /app/app.env
/app/migrate -path /app/migration -database "" -verbose up

echo "start the app"
exec "$@"$DB_SOURCE