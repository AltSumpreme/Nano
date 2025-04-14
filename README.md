
# Nano 

**Nano** is a scalable microservices-based backend for a blogging platform, written in **Go**, using **gRPC**, **GraphQL**, **PostgreSQL**, **Elasticsearch**, and **Docker**.

---

> [!WARNING]
> 
> Nano is a WIP
> 
> Most aspects of the project are under heavy development
> and no stable release is present as of yet.
> 

##  Architecture Overview

Nano is composed of the following independent services:

| Service         | Purpose                        | Database        | Communication |
|----------------|--------------------------------|-----------------|---------------|
| Auth Service    | Authentication & JWT issuance  | PostgreSQL      | gRPC          |
| Post Service    | Post creation & management     | PostgreSQL      | gRPC          |
| Search Service  | Full-text blog post search     | Elasticsearch   | gRPC          |
| GraphQL Gateway | Unified API interface          | —               | GraphQL + gRPC|

Each service runs in its own container and can be scaled independently.

---

## Project Structure

```
nano/
├── auth-service/           # User authentication and token service
│   ├── cmd/
|   |     ├──account       
|   |            ├──main.proto
|   ├──pb/                 # generated proto files
|   ├── db.dockerfile
|   ├── app.dockerfile
|   ├── account.sql
|   ├── client.go
|   ├── server.go
│   ├── repository.go          
│   └── service.go
│
├── post-service/           # Handles post creation, retrieval, editing
│   ├── proto/
│   ├── internal/
│   └── main.go
│
├── search-service/         # Uses Elasticsearch for full-text search
│   ├── cmd/
|   |     ├──account       
|   |            ├──main.proto
|   |
|   ├──pb/                 # generated proto files
|   ├── app.dockerfile
|   ├── client.go
|   ├── server.go
│   ├── repository.go          
│   └── service.go
├── graphql-api-gateway/    # API Gateway exposing GraphQL API
│   ├── account_resolver
│   ├── gqlgen.yaml
│   └── main.go
|   ├── generated.go
|   ├── models.go
|   ├──model_gen.go
|   ├── mutation_resolver.go
|   ├── query_resolver
|   ├──schmea.graphql
│
│
└── docker-compose.yml      # Service orchestration
├── go.mod
├── go.sum
```

---

##  Tech Stack

- **Go**: Core programming language
- **gRPC**: Service-to-service communication
- **GraphQL**: Unified API exposed to frontend
- **PostgreSQL**: Database for Auth & Post services
- **Elasticsearch**: High-performance full-text search
- **Docker**: Containerization of all services
- **(Future)**: ELK Stack (Elasticsearch, Logstash, Kibana) for observability

---

##  Communication Flow

```plaintext
Frontend → GraphQL Gateway → [Auth | Post | Search] Services via gRPC
```

- **Auth**: Validates users and tokens.
- **Post**: Manages blog posts (currently WIP).
- **Search**: Queries posts from Elasticsearch.
- **GraphQL Gateway**: Aggregates all service endpoints into a unified API.

---

##  Setup & Run (Dev)

### 1. Prerequisites
- Go 1.21+
- Docker & Docker Compose
- Protobuf compiler

### 2. Clone the repo
```bash
git clone https://github.com/yourname/nano.git
cd nano
```

### 3. Build & Run all services
```bash
docker-compose up --build
```

---

##  Future Plans

- [ ] Complete Post Service
- [ ] Implement ELK stack for logging and observability
- [ ] Add rate limiting and circuit breaking
- [ ] Add CI/CD with GitHub Actions

---

## Contributing

Contributions are welcome! Feel free to open issues or pull requests.

---

## License

MIT © 2025 
