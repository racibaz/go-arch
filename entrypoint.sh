#!/bin/bash
set -e

echo "Running database migrations..."
make db_migrate_up

#echo "Seeding database..."
make seed

#echo "Starting server with Air..."
air