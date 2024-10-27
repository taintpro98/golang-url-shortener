#!/bin/sh
echo "Running database migrations..."
/bin/migration -dir /app/migrations up

echo "Starting the server..."
exec "$@"