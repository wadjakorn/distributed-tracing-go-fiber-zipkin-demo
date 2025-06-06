# Distributed Tracing Demo
## Using Go Fiber Zipkin

```
go run .zipkin-demo-service-1/...

go run .zipkin-demo-service-2/...
```

### Test Get
```
curl --location 'http://localhost:4000/orders'
```

### Test Post
```
curl --location 'http://localhost:4000/order' \
--header 'Content-Type: application/json' \
--data '{
    "amount": 0,
    "status": "Pending"
}'
```