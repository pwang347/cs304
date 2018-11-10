#!/bin/bash

MYSQL_USER=${MYSQL_USER:-root}

# Create database and tables and then insert fake data
if [ -z $MYSQL_PASSWORD ]; then
    mysql -u $MYSQL_USER < bootstrap-db.sql
    go run seed.go -user $MYSQL_USER
else
    mysql -u $MYSQL_USER -p $MYSQL_PASSWORD < bootstrap-db.sql
    go run seed.go -user $MYSQL_USER -password $MYSQL_PASSWORD
fi
