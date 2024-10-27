#!/bin/sh
echo "Running database migrations..."
echo $POSTGRES_HOST
echo $POSTGRES_DB
echo $POSTGRES_USER
/bin/migration -dir /app/migrations up

echo "Starting the server..."
exec "$@"