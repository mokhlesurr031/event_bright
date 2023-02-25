# event_bright


## API Endpoints 


### Users
1. Registration
    - **Request**
    ```
   URL: {{domain}}/api/v1/auth/
   Method: POST
   Payload:
    {
    "name": "Test User",
    "email":"test@gmail.com",
    "password":"1234",
    "password_confirm":"1234"
    }
    ```
    - **Response:**
   ```
   Status Code: 201
   {
   "id": 6,
   "name": "Test User 4",
   "email": "test4@gmail.com",
   "password": "",
   "password_confirm": "",
   "created_at": "2023-02-25T06:31:42.090517+06:00"
   }
    
    ```


2. Login
   - **Request**
    ```
   URL: {{domain}}/api/v1/auth/login/
   Method: POST
   Payload:
    {
    "email": "test@gmail.com",
    "password": "1234"
   }
    ```
   - **Response:**
   ```
   {
       "ExpiredIn": 3600000000000,
       "MaxAge": 3600,
       "Secret": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzcyODYwMzUsImlkIjoiMiJ9.2PPQCrvb6mnB_S3KZRQTmWfjEvWzlTXBVojKcXGjzGw",
       "Message": "success",
       "User": {
         "id": 2,
         "name": "Test User",
         "email": "test@gmail.com"
       }
   }
    
    ```

### Events

1. Create Event
   - **Request**
    ```
   URL: {{domain}}/api/v1/event/
   Method: POST
   Authorization: Bearer Token
   Payload:
   {
   "name": "Test Event",
   "location": "Dhaka",
   "description": "Descc",
   "date": "2023-02-25 23:56" //24 hours time format
   }
    ```
   - **Response:**
   ```
   Status Code: 201
   {
   "id": 21,
   "name": "Test Event",
   "date": "2023-02-25T23:56:00Z",
   "location": "Dhaka",
   "description": "Descc",
   "created_by": 2,
   "total_participant": 0,
   "created_at": "2023-02-26T01:17:59.407923+06:00"
   }
    
    ```

2. List
   - **Request**
    ```
   URL: {{domain}}/api/v1/event/list
   Method: GET
   Payload: N/A
    ```
   - **Response:**
   ```
   Status Code: 200
   [
   {
   "id": 1,
   "name": "Test Event",
   "date": "0001-01-01T06:01:40+06:01",
   "location": "Dhaka",
   "description": "Descc",
   "created_by": 1,
   "total_participant": 0,
   "created_at": "2023-02-22T14:02:29.163401+06:00"
   },
   {
   "id": 2,
   "name": "Test Event",
   "date": "0001-01-01T06:01:40+06:01",
   "location": "Dhaka",
   "description": "Descc",
   "created_by": 1,
   "total_participant": 0,
   "created_at": "2023-02-22T14:03:49.645269+06:00"
   },
   {
   "id": 3,
   "name": "Test Event",
   "date": "0001-01-01T06:01:40+06:01",
   "location": "Dhaka",
   "description": "Descc",
   "created_by": 1,
   "total_participant": 0,
   "created_at": "2023-02-22T14:04:22.433178+06:00"
   },
   {
   "id": 4,
   "name": "Test Event",
   "date": "0001-01-01T06:01:40+06:01",
   "location": "Dhaka",
   "description": "Descc",
   "created_by": 1,
   "total_participant": 17,
   "created_at": "2023-02-22T14:06:37.526819+06:00"
   },
   {
   "id": 5,
   "name": "Test Event",
   "date": "0001-01-01T06:01:40+06:01",
   "location": "Dhaka",
   "description": "Descc",
   "created_by": 2,
   "total_participant": 0,
   "created_at": "2023-02-25T05:56:50.487284+06:00"
   }
   ]
    
    ```

3. My Event List
   - **Request**
    ```
   URL: {{domain}}/api/v1/event/list/me
   Method: GET
   Authorization: Bearer Token
   Payload: N/A

    ```
   - **Response:**
   ```
   Status Code: 200
   [
   {
   "id": 5,
   "name": "Test Event",
   "date": "0001-01-01T06:01:40+06:01",
   "location": "Dhaka",
   "description": "Descc",
   "created_by": 2,
   "total_participant": 0,
   "created_at": "2023-02-25T05:56:50.487284+06:00"
   }
   ]
    
    ```

4. Event Details
   - **Request**
    ```
   URL: {{domain}}/api/v1/event/list/{{event_id}}
   Method: GET
   Payload: N/A
    ```
   - **Response:**
   ```
   Status Code: 200
   {
   "id": 1,
   "name": "Test Event",
   "date": "0001-01-01T06:01:40+06:01",
   "location": "Dhaka",
   "description": "Descc",
   "created_by": 1,
   "total_participant": 0,
   "created_at": "2023-02-22T14:02:29.163401+06:00"
   }
    
    ```

5. Participate Event
   - **Request**
    ```
   URL: {{domain}}/api/v1/event/list/{{event_id}}/go
   Method: POST
   Payload:
   {
   "name": "JEFFY",
   "email": "jeffy2@gmail.com"
   }
    ```
   - **Response:**
   ```
   Status Code: 201
    {
    }
    
    ```
