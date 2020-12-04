#!/bin/bash

# Create iris table
docker container exec -it mlgo_postgres psql -U postgres -c \
    "CREATE TABLE iris ( \
        sepal_length FLOAT, \
        sepal_width FLOAT, \
        petal_length FLOAT, \
        petal_width FLOAT, \
        species VARCHAR(20) \
    );"

# Put iris.csv to the table
docker container exec -it mlgo_postgres psql -U postgres -c \
    "COPY iris(sepal_length, sepal_width, petal_length, petal_width, species) \
     FROM '/storage/data/iris.csv'
     DELIMITER ','
     CSV HEADER;"
