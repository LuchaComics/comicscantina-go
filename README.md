# comicscantina-go
The web platform for buying and selling comic books online. The ``go`` repository is the back-end service which powers a front-end service.

## Installation
1. Get the latest code.

  ```bash
  go get -u -v github.com/luchacomics/comicscantina-go
  ```


2. Open up postgres and run the following to setup the database for development or production.

  ```sql
  drop database comicscantina_db;
  create database comicscantina_db;
  \c comicscantina_db;
  CREATE USER golang WITH PASSWORD 'YOUR_PASSWORD';
  GRANT ALL PRIVILEGES ON DATABASE comicscantina_db to golang;
  ALTER USER golang CREATEDB;
  ALTER ROLE golang SUPERUSER;
  CREATE EXTENSION postgis;
  ```


3. Run the following to setup the database for unit testing.

  ```sql
  drop database comicscantina_test_db;
  create database comicscantina_test_db;
  \c comicscantina_test_db;
  CREATE USER golang WITH PASSWORD 'YOUR_PASSWORD';
  GRANT ALL PRIVILEGES ON DATABASE comicscantina_test_db to golang;
  ALTER USER golang CREATEDB;
  ALTER ROLE golang SUPERUSER;
  CREATE EXTENSION postgis;
  ```


4. Install our dependencies

  ```
  ./requirements.sh
  ```


5. Update environmental variables by running the following. Please change the variables to whatever you prefer.

  ```bash
  #!/bin/bash
  export COMICSCANTINA_GORM_CONFIG="postgres://golang:YOUR_PASSWORD@localhost/comicscantina_db?sslmode=disable"
  export COMICSCANTINA_SECRET="YOUR_SECRET_RANDOM_STRING"
  export COMICSCANTINA_ADDRESS="127.0.0.1:8080"
  export COMICSCANTINA_UNIT_TEST_GORM_CONFIG="postgres://golang:YOUR_PASSWORD@localhost/comicscantina_test_db?sslmode=disable"
  ```
