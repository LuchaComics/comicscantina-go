Comics Cantina API Reference
======

## Developers Note

To help make the next few API endpoints easy to type, save your token to the console.

```
COMICS_WS_API_TOKEN='YOUR_TOKEN'
```


## Get API Version
Returns the version information.

* **URL**

  `/api/v1/public/version`


* **Method**

  `GET`


* **URL Params**

  None


* **Data Params**

  None


* **Success Response**

  * **Code:** 200
  * **Content:** ``{"Service": `"v0.1", "API: 'v1'"}``


* **Error Response**

  * None


* **Sample Call**

  ``
  $ http get 127.0.0.1:8080/api/v1/public/version
  ``


## Register
Registers the user account with our system.

* **URL**

  ``/api/v1/public/register``


* **Method**

  `POST`


* **URL Params**

  None


* **Data Params**

  * email
  * password
  * first_name
  * last_name


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```
    {
        "email": "bart@mikasoftware.com",
        "first_name": "Bart",
        "last_name": "Mika",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDkyOTY1MDAsInVzZXJfaWQiOjF9.QN9dyWL2dlxKgkm0xbQAmnaI6_4amHcSfqUGQ6pZbxM",
        "user_id": 1
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```
    {
        "error": "Email is not unique.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```
  $ http post 127.0.0.1:8080/api/v1/public/register email=bart@mikasoftware.com password=123password first_name=Bart last_name=Mika
  ```


## Login
Returns the user profile and authentication token upon successful login in.

* **URL**

  ``/api/v1/public/login``


* **Method**

  `POST`


* **URL Params**

  None


* **Data Params**

  * email
  * password


* **Success Response**

  * **Code:** 200
  * **Content:**

    ```
    {
        "email": "bart@mikasoftware.com",
        "first_name": "Bart",
        "last_name": "Mika",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDkyOTg1MDYsInVzZXJfaWQiOjF9.HrwHvfL4-1pMe7EcXEzlsxciFgK0xf2uC8BV1kfLT_c",
        "user_id": 1
    }
    ```


* **Error Response**

  * **Code:** 400
  * **Content:**

    ```
    {
        "error": "Email or password is incorrect.",
        "status": "Invalid request."
    }
    ```


* **Sample Call**

  ```
  $ http post 127.0.0.1:8080/api/v1/public/login email=bart@mikasoftware.com password=123password
  ```


## Get Profile
The API endpoint used to get the user profile. Only the profile of the
authenticated user is returned.

* **URL**

  ``/api/v1/profile``


* **Method**

  ``GET``


* **URL Params**

    None


* **Data Params**

    None


* **Success Response**

    * **Code:** 200
    * **Content:**

    ```
    {
        "email": "bart@mikasoftware.com",
        "first_name": "Bart",
        "last_name": "Mika",
        "user_id": 1
    }
    ```


* **Error Response**

    * None


* **Sample Call**

    ```
    $ http get 127.0.0.1:8080/api/v1/profile Authorization:"Bearer $COMICS_WS_API_TOKEN"
    ```

# ---------------------------------------
# ---------------------------------------
# ---------------------------------------
# ---------------------------------------



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

## Organization - Create
### Description

  ```
  /api/v1/organizations
  ```

The API endpoint used to create the organization.

### Example Command

  ```
  http post 127.0.0.1:8080/api/v1/stores Authorization:"Bearer $COMICS_WS_API_TOKEN" name="Main Store" description="The brick and morter comics store." email="bart@mikasoftware.com" street_address="111-204 Infinite Loop Road" city="London" province="Ontario" country="Canada" organization_id=1
  ```

### Example Output

  ```
  ```
