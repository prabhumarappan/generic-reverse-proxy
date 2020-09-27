# freshworks-hiring

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
