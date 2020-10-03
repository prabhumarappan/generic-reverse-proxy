# reverse-proxy

Request Proxy Service

Language : Go

Details:
========
1. Forward request as per the URL in the request body
2. Reject requests with non-https
3. Timeout requests which take more than 5 seconds to respond


Middlewares Used:
=================
1. Request Validator
2. Logging Middleware
3. Rate Limiter

Build
=====
1. Go to cmd/proxy-server
2. Run go build main.go

Deployment
==========
1. Go to cmd/proxy-server
2. Run ./main
3. Or Run go run main.go

Test
=====
1. On the root folder of the project
2. Run go test ./...

Request Sample
==============
```curl
curl --location --request GET 'localhost:8083/proxy' \
--header 'Content-Type: application/json' \
--data-raw '{
    "URL": "https://duckduckgo.com",
    "RequestType": "GET", 
    "ClientID": "a1ffdc3b-0204-4410-816c-5edaccc4eaad",
    "Headers": {
        "Name": "Prabhu",
        "Gender": "Male"
    },
    "RequestBody": "Hello World"
}'
```

Rate Limiter
============
I was initially thinking of going with a simple Rate Limiter based on the IP address, but since we have the ClientId with us, I then pivoted to the idea that we can rate limit on the basis of the ClientId which will be more usefull from the information that we have!

Additionally, apart from the ClientId rate limiting, we can also have another rate limiter on the basis of IP address too.
