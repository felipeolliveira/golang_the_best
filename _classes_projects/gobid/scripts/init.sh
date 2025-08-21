#!/bin/bash

POSTGRES_URL=postgresql://$GOBID_DATABASE_USER:$GOBID_DATABASE_PASSWORD@$GOBID_DATABASE_HOST:$GOBID_DATABASE_PORT/$GOBID_DATABASE_NAME

echo "Waiting for database to be ready..."
until pg_isready -d "$POSTGRES_URL"; do
  sleep 1
done

echo "Creating database if doesn't exists..."
psql $POSTGRES_URL -c "CREATE DATABASE $GOBID_DATABASE_NAME;" || true

echo "Applying migrations..."
cd /app/internal/store/pgstore/migrations
tern migrate

echo "Starting application..."
cd /app
exec ./main
