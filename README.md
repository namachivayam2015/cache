# Cache Application

This project implements basic capabilities of Cache using Golang v1.13.4

Implemented Unit Test for util package & validate REST Endpoints using httptest package

Implemented a goroutine to remove least used key based on the readThreshold loaded from environment variable.


# Endpoints Exposed:

*Port :8080*

#### Headers


*Content-Type* : application/json

*Authorization* : Basic base64($BASIC_AUTH_USERNAME:$BASIC_AUTH_PASSWORD)

|Method|URI|Description|
|------|----|----------|
|POST|/add|Add Key/Value to Cache|
|PUT|/update|Update Value of a given Key in Cache|
|DELETE|/remove/{key}|Remove Key from Cache|
|GET|/fetch/{key}|Retrieve the Value of a given Key|
|GET|/fetchall|Retrieve all Key Value pairs in the Cache|



# To run the application:

go run app.go

# To build image & run the image


cd src/

docker build . -t go-cache

docker run --env-file .\variables.env -d -p 8080:8080 go-cache:latest


##### Environment Variables required to run the application:

- CACHE_GC_INTERVAL=5
- DEFAULT_CACHE_SIZE=25
- CACHE_READ_TIMEOUT_THRESHOLD=30 
- BASIC_AUTH_USERNAME=testuser
- BASIC_AUTH_PASSWORD=testpassword
