#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $POSTGRES_USER WITH PASSWORD '$POSTGRES_PASS';
  CREATE DATABASE $POSTGRES_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_NAME TO $POSTGRES_USER;
  \connect $POSTGRES_NAME $POSTGRES_USER
  COMMIT;
  EOSQL