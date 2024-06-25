#!/bin/bash

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=$POSTGRES_PASSWORD psql -h "$host" -U "postgres" -lqt | cut -d \| -f 1 | grep -qw "sugar-db"; do
  echo "Waiting for PostgreSQL to be ready..."
  sleep 1
done

echo "PostgreSQL is ready - executing command"
exec $cmd