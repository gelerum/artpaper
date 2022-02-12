FROM postgres:14.1-alpine3.15

WORKDIR /sql

COPY ./init.sql /docker-entrypoint-initdb.d/
