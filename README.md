# Distributed Tracing Demo
## Using Go Fiber Zipkin

```
docker compose up -d

// or

podman compose up -d

cd zipkin-demo-service-1
go run .

cd zipkin-demo-service-2
go run .
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