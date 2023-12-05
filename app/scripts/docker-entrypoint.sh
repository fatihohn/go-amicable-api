#!/bin/bash
set -e

cp "/app/db/schema.sql" "/app/db/schema.sql.bak"
sed -i "s/POSTGRES_DB/$POSTGRES_DB/g" /app/db/schema.sql
sed -i "s/POSTGRES_USER/$POSTGRES_USER/g" /app/db/schema.sql
sed -i "s/POSTGRES_PASSWORD/$POSTGRES_PASSWORD/g" /app/db/schema.sql

# Wait for the database to be ready (replace with your actual database readiness check)
until pg_isready -h $POSTGRES_HOST -p $POSTGRES_PORT -U $POSTGRES_USER -d $POSTGRES_DB; do
  echo "Waiting for the database to be ready..."
  sleep 1
done

# Run your schema.sql script (replace with your actual script and database client command)
PGPASSWORD=$POSTGRES_PASSWORD psql -h $POSTGRES_HOST -p $POSTGRES_PORT -U $POSTGRES_USER -d $POSTGRES_DB -a -f ./db/schema.sql

unset PGPASSWORD

# Run the main Go application
./main
