#!/bin/bash

echo "Starting Migration"

node './styles/JsonToSQL.js'

node './questions/JsonToSQL.js'

psql -f delete_tables.sql bjcp_table_dev

psql -f create_db.sql

psql -f create_tables.sql bjcp_table_dev

psql -f ./questions/questions-db.sql bjcp_table_dev

psql -f ./styles/styles-db.sql bjcp_table_dev

rm ./questions/questions-db.sql

rm ./styles/styles-db.sql

echo "Migration Completed"
