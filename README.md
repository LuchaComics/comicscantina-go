# comicscantina-go
The web platform for buying and selling comic books online. The ``go`` repository is the back-end service which powers a front-end service.

## Installation

1. Open up postgres and run the following to setup the database for development or production.

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

2. Run the following to setup the database for unit testing.

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

3. Install our dependencies

  ```
  ./requirements.sh
  ```

4. Update environmental variables by running the following. Please change the variables to whatever you prefer.

  ```bash
  #!/bin/bash
  export COMICSCANTINA_GORM_CONFIG="host=localhost port=5432 user=golang dbname=comicscantina_db password=YOUR_PASSWORD sslmode=disable"
  export COMICSCANTINA_SECRET="YOUR_SECRET_RANDOM_STRING"
  export COMICSCANTINA_ADDRESS="127.0.0.1:8080"
  export COMICSCANTINA_UNIT_TEST_GORM_CONFIG="host=localhost port=5432 user=golang dbname=comicscantina_test_db password=YOUR_PASSWORD sslmode=disable"
  ```
