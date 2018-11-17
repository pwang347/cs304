#!/bin/bash

MYSQL_USER=${MYSQL_USER:-root}
MYSQL_HOST=${MYSQL_HOST:-localhost}

# Create database and tables and then insert fake data
if [ -z $MYSQL_PASSWORD ]; then
    mysql -u $MYSQL_USER -h $MYSQL_HOST < bootstrap-db.sql
    go run seed.go -user $MYSQL_USER -mysqlHost $MYSQL_HOST
else
    mysql -u $MYSQL_USER -p $MYSQL_PASSWORD -h $MYSQL_HOST < bootstrap-db.sql
    go run seed.go -user $MYSQL_USER -password $MYSQL_PASSWORD -mysqlHost $MYSQL_HOST
fi
