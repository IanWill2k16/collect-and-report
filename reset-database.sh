#! /bin/bash

goose -dir ./sql/schema postgres "postgres://postgres:postgres@localhost:5432/collectandreport" down-to 0
goose -dir ./sql/schema postgres "postgres://postgres:postgres@localhost:5432/collectandreport" up