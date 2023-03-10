### Events

1. Create Event
    - **Request**
    ```
   URL: {{domain}}/api/v1/event/
   Method: POST
   Authorization: Bearer Token
   Payload:
   {
     "name": "Test Event-6",
     "location": "Khulna",
     "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
     "date": "2023-03-05 15:00"
   }
    ```
    - **Response:**
   ```
   Status Code: 201
   {
     "id": 6,
     "name": "Test Event-6",
     "date": "2023-03-05T15:00:00Z",
     "location": "Khulna",
     "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
     "created_by": 3,
     "total_participant": 0,
     "created_at": "2023-02-26T01:39:33.210477+06:00"
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
        "id": 2,
        "name": "Test Event-2",
        "date": "2023-03-05T21:00:00+06:00",
        "location": "Chittagong",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "created_by": 1,
        "total_participant": 0,
        "created_at": "2023-02-26T01:37:03.105009+06:00"
    },
    {
        "id": 1,
        "name": "Test Event",
        "date": "2023-03-05T21:00:00+06:00",
        "location": "Dhaka",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "created_by": 1,
        "total_participant": 0,
        "created_at": "2023-02-26T01:36:14.173524+06:00"
    },
    {
        "id": 3,
        "name": "Test Event-3",
        "date": "2023-03-05T21:00:00+06:00",
        "location": "Rajshahi",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "created_by": 2,
        "total_participant": 0,
        "created_at": "2023-02-26T01:38:33.056597+06:00"
    },
    {
        "id": 4,
        "name": "Test Event-4",
        "date": "2023-03-05T21:00:00+06:00",
        "location": "Barishal",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "created_by": 2,
        "total_participant": 0,
        "created_at": "2023-02-26T01:38:39.674905+06:00"
    },
    {
        "id": 5,
        "name": "Test Event-5",
        "date": "2023-03-05T21:00:00+06:00",
        "location": "Sylhet",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "created_by": 3,
        "total_participant": 0,
        "created_at": "2023-02-26T01:39:09.860104+06:00"
    },
    {
        "id": 6,
        "name": "Test Event-6",
        "date": "2023-03-05T21:00:00+06:00",
        "location": "Khulna",
        "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
        "created_by": 3,
        "total_participant": 0,
        "created_at": "2023-02-26T01:39:33.210477+06:00"
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
       "name": "Test Event-5",
       "location": "Khulna",
       "date": "2023-03-05T21:00:00+06:00",
       "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
       "created_by": "Test User-3",
       "total_participant": 0,
       "Participant": []
     },
     {
       "name": "Test Event-5",
       "location": "Sylhet",
       "date": "2023-03-05T21:00:00+06:00",
       "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
       "created_by": "Test User-3",
       "total_participant": 4,
       "Participant": [
       {
         "name": "Participant User 1",
         "email": "par2@gmail.com"
       },
       {
         "name": "Participant User 2",
         "email": "par1@gmail.com"
       },
       {
         "name": "Participant User 3",
         "email": "par3@gmail.com"
       },
       {
         "name": "Participant User 4",
         "email": "par4@gmail.com"
       }
      ]
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
     "date": "2023-03-05T21:00:00+06:00",
     "location": "Dhaka",
     "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
     "created_by": 1,
     "total_participant": 0,
     "created_at": "2023-02-26T01:36:14.173524+06:00"
   }
    
    ```

5. Participate Event
    - **Request**
    ```
   URL: {{domain}}/api/v1/event/list/{{event_id}}/go
   Method: POST
   Payload:
   {
     "name": "Participant User 1",
     "email": "par1@gmail.com"
   }
    ```
    - **Response:**
   ```
   Status Code: 201
   {
     "id": 2,
     "name": "Participant User 1",
     "email": "par1@gmail.com",
     "event_id": 1,
     "created_at": "2023-02-26T01:44:33.447431+06:00"
   }
   
   If this user has already participated in this event:
   
   Status: 500
   Error: You have already participated in this event
    
    ```

6. Participant Count Updated After Participating a particular event

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
     "date": "2023-03-05T21:00:00+06:00",
     "location": "Dhaka",
     "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
     "created_by": 1,
     "total_participant": 2,  //Participant Count
     "created_at": "2023-02-26T01:36:14.173524+06:00"
   }
    
    ```
