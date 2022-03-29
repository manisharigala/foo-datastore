# foo-datastore

## Goal
To build a very simple webservice in go that will store “foo” records in memory

## Objectives / Assumptions 
Below mentioned points pertain to the certain assumptions and/or constraints:
- Impelemt in memory datastore
- Webservice should run on port 8080.
- The foo data structure only needs 2 fields. A “name” field and an “id” field. Both should be string data types. 
- POST endpoint (‘/foo’) that accepts a json foo object (with just "name" field) and responds with a 200 response code. The value of the id field should be added by this endpoint using a generated UUID.
- If another json foo object with same name is requested with POST endpoint, new UUID is generated and stored as a sepetate record.
- Support a GET endpoint (‘foo/{id}’) that responds with a 200 response code if the record is found, or a 404 response code if not found
- Support a DELETE endpoint (‘foo/{id}’) that responds with a 204 response code if the record is found, or a 404 response code if not found.
- Code is formatted with gofmt.
- Include a Makefile that supports “build”, “run”, and “clean” tasks.

## Design and Implementation
**HTTP Server** : The HTTP REST Server is built using gorilla/mux. Clients can request GET/POST/DELETE APIs from the REST Server. Based on request type, the server responds. Failure and exception handling is implemented.
**In-Memory Storage**: Data is stored using a hash map for quick data accesss.

## Dependencies
| Dependency    | Version | Link  |
| ------------- |-------- |-----:|
| Go            | 1.16    | https://go.dev/ |
| UUID          | 1.3.0   |https://github.com/google/uuid |
| gorilla/mux  |  1.8.0   |github.com/gorilla/mux  |

## API Documention

- **GET /**
    - Ping check to check REST Server is up
    - Request
        None
    - Response
        msg: String

- **POST /foo**
    - Upload a json file to server
    - Request
        name : String
    - Response
        status : Number [200 OK, 400 Bad Request]
        { name : String
          uuid : String }

- **GET /foo{id}**
    - Get foo object associated with id
    - Request 
        id : String
    - Response
        status : Number [200 OK, 404 Not Found]
        { name : String
          uuid : String }
- **DELETE /foo{id}**
    - Delete foo object associated with id from memeory
    - Request
        id : String
    - Response
        Status : Number [204 No Content, 404 Not Found]
        

## Instructions to Invoke Application
- Clone the github repository.
- Open terminal in the root folder where the  makefile is present and run either of the following commands.
    - make run
        - To built and run the application
    - make build
        - To build the executable file that is store in bin/main
3. Sample Curl operations
    - $ curl -i -X POST -H 'Accept: application/json' -H 'Content-Type: application/json' 'http://localhost:8080/foo' -d '{"name": "Jack"}'
    - $ curl -i -X GET -H 'Accept: application/json' 'http://localhost:8080/foo/26baf48a-db0f-4884-9b89-820ce7596a6e'
    - $ curl -i -X DELETE 'http://localhost:8080/foo/26baf48a-db0f-4884-9b89-820ce7596a6e'




