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
  HTTP/1.1 200 OK
  Content-Length: 73
  Content-Type: application/json
  Date: Fri, 01 Feb 2019 01:21:46 GMT

  {
      "Details": "Welcome to the ComicsCantina backend API, build v0.0.001.0"
  }
  ```

## Register
### Description
Registers the user account with our system.

  ```
  /api/v1/register
  ````

### Example Command

  ```
  http post 127.0.0.1:8080/api/v1/register email=bart@mikasoftware.com password=123password first_name=Bart last_name=Mika
  ```

### Example Output

  ```
  HTTP/1.1 201 Created
  Content-Length: 134
  Content-Type: application/json; charset=utf-8
  Date: Fri, 01 Feb 2019 01:22:15 GMT

  {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDkwNzA1MzUsInVzZXJfaWQiOjF9.S13lwn4lSHMNCSi2D0bbDNjNNKcjnTLGvBrzbFWNZDw"
  }
  ```

## Login
### Description
Gets a new token based on the ``email`` and ``password``.

  ```
  /api/v1/login
  ```

### Example Command

  ```
  http POST 127.0.0.1:8080/api/v1/login email=bart@mikasoftware.com password=123password
  ```

### Example Output

  ```
  HTTP/1.1 200 OK
  Content-Length: 134
  Content-Type: application/json; charset=utf-8
  Date: Fri, 01 Feb 2019 01:24:25 GMT

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

The API endpoint used to get the user profile.

### Example Command

  ```
  http get 127.0.0.1:8080/api/v1/profile Authorization:"Bearer $COMICS_WS_API_TOKEN"
  ```

### Example Output

  ```
  HTTP/1.1 200 OK
  Content-Length: 46
  Content-Type: application/json; charset=utf-8
  Date: Fri, 01 Feb 2019 14:09:35 GMT

  {
      "email": "bart@mikasoftware.com",
      "user_id": 1,
      "first_name": "Bart",
      "last_name": "Mika"
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
  http post 127.0.0.1:8080/api/v1/organizations Authorization:"Bearer $COMICS_WS_API_TOKEN" name="Mika Software" description="The company" email="bart@mikasoftware.com"
  ```

### Example Output

  ```
  HTTP/1.1 201 Created
  Content-Length: 105
  Content-Type: application/json; charset=utf-8
  Date: Fri, 01 Feb 2019 18:39:11 GMT

  {
      "description": "The company",
      "email": "bart@mikasoftware.com",
      "id": 3,
      "name": "Mika Software",
      "owner_id": 1
  }
  ```
