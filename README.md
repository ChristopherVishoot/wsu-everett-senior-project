# WSU Senior Project

A Go microservice built with [Go-kit](https://gokit.io/) following a layered architecture pattern, deployed on Kubernetes.

## Architecture

The service follows Go-kit's standard layered architecture:

- **Service Layer** (`golang/service/`) - Business logic and database operations
- **Endpoint Layer** (`golang/endpoint/`) - Request/response types and endpoint definitions
- **Transport Layer** (`golang/transport/`) - HTTP handlers and JSON encoding/decoding
- **Database Layer** (`golang/database/`) - PostgreSQL connection and configuration

## Running Locally

```bash
cd golang
./compile.sh
```

## Kubernetes Deployment

Apply all infrastructure manifests:

```bash
kubectl apply -f infrastructure/ --recursive
```

### Port Forward the Service

To access the service locally:

```bash
kubectl port-forward svc/golang-service 8080:8080
```

## API Examples

### Create a User

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

### Get All Users

```bash
curl http://localhost:8080/users/all
```
