# API Reference
## Health Status API

### URL
### Description
Returns a status of the API endpoint.

  ```
  /
  ```

### Example Command

  ```
  http get 127.0.0.1:8080/
  ```

### Example Output

  ```
  {
      "Details": "Welcome to the ComicsCantina backend API, build v0.0.001.0"
  }
  ```

## Register
### Description
Registers the user account with our system.

  ```
  /api/v1/public/register
  ````

### Example Command

  ```
  http post 127.0.0.1:8080/api/v1/public/register email=bart@mikasoftware.com password=123password first_name=Bart last_name=Mika
  ```

### Example Output

  ```
  {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDkwNzA1MzUsInVzZXJfaWQiOjF9.S13lwn4lSHMNCSi2D0bbDNjNNKcjnTLGvBrzbFWNZDw"
  }
  ```

### Notes:
To help make the next few API endpoints easy to type, save your token to the console.

  ```
  COMICS_WS_API_TOKEN='YOUR_TOKEN'
  ```

## Login
### Description
Gets a new token based on the ``email`` and ``password``.

  ```
  /api/v1/public/login
  ```

### Example Command

  ```
  http POST 127.0.0.1:8080/api/v1/public/login email=bart@mikasoftware.com password=123password
  ```

### Example Output

  ```
  {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDkwNzA2NjUsInVzZXJfaWQiOjF9.s16HtRf7qzu5vDV1iY3c5g8UMYfS1cvM3bV9e3ammxA"
  }
  ```

### Notes:
To help make the next few API endpoints easy to type, save your token to the console.

  ```
  COMICS_WS_API_TOKEN='YOUR_TOKEN'
  ```

## Profile
### Description

  ```
  /api/v1/profile
  ```

The API endpoint used to get the user profile. Only the profile of the
authenticated user is returned.

### Example Command

  ```
  http get 127.0.0.1:8080/api/v1/profile Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```

### Example Output

  ```
  {
      "email": "bart@mikasoftware.com",
      "user_id": 1,
      "first_name": "Bart",
      "last_name": "Mika"
  }
  ```

## Organization - List
### Description

  ```
  /api/v1/organizations
  ```

The API endpoint used to list the organizations.

### Example Command

  ```
  http get 127.0.0.1:8080/api/v1/organizations Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```

### Example Output

  ```
  {
      "description": "The company",
      "email": "bart@mikasoftware.com",
      "id": 3,
      "name": "Mika Software",
      "owner_id": 1
  }
  ```

## Organization - Create
### Description

  ```
  /api/v1/organizations
  ```

The API endpoint used to create the organization.

### Example Command

  ```
  http post 127.0.0.1:8080/api/v1/organizations Authorization:"Bearer $COMICS_WS_API_TOKEN" name="Mika Software" description="The company" email="bart@mikasoftware.com" street_address="111-204 Infinite Loop Road" city="London" province="Ontario" country="Canada"
  ```

### Example Output

  ```
  {
      "Status": 0,
      "city": "London",
      "country": "Canada",
      "description": "The company",
      "email": "bart@mikasoftware.com",
      "id": 1,
      "name": "Mika Software",
      "owner_id": 1,
      "province": "Ontario",
      "street_address": "111-204 Infinite Loop Road"
  }
  ```



## Stores - List
### Description

  ```
  /api/v1/stores
  ```

The API endpoint used to list the stores in the system.

### Example Command

  ```
  http get 127.0.0.1:8080/api/v1/stores Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```

### Example Output

  ```
  {
    "description": "The company",
    "email": "bart@mikasoftware.com",
    "id": 3,
    "name": "Mika Software",
    "owner_id": 1
  }
  ```
