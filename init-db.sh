#!/bin/bash
set -e

echo "Starting database creation..."

# Buat beberapa database
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="postgres" <<-EOSQL
    CREATE DATABASE "braciate-prod";
    CREATE DATABASE "braciate-staging";
EOSQL

echo "Databases created successfully!"
