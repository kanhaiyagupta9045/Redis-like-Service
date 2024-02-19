### Redis-like Service

### Running the Server
To start the Redis-like server, execute the following command in your terminal:

```
$ go run main.go
```
### Using the Service
```
STEP 1 : Make a POST request to the following URL: http://localhost:8080/operation
STEP 2 : Include JSON data in your POST request 
STEP 3 : like this {
  "operation": "SET key_b 3 EX 60 XX"
}
STEP 4 : to test Other Commands pass json data like this
{
  "operation": "GET key_b"
}

```
# Working
```
When making a POST request, the JSON data is first validated. Then, depending on the specific command, the request is processed accordingly
```
