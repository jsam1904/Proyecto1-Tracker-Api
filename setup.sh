#!/bin/bash
set -e

echo "Installing Go dependencies..."
go mod tidy

echo "Creating database..."
psql -U postgres -c "CREATE DATABASE bundesliga;" 2>/dev/null || echo "Database may already exist"

echo "Running migrations..."
psql -U postgres -d bundesliga -f db/migrations/001_init.sql

echo "Done! Run: go run main.go"