# 🐦 TwitSocial – Social Media REST API

TwitSocial is a backend project developed using **Go (Golang) and the Gin Web Framework**, structured to support a Twitter-like social networking application.

The project focuses on RESTful API development, JWT-based authentication, relational database integration, schema migrations, and modular backend architecture.

It uses **MySQL 8.0**, Docker Compose for database provisioning, and dbmate for version-controlled database migrations.

## 🚀 Key Features

### 🔐 Authentication & Security

- JWT libraries for token-based authentication.
- Cryptographic utilities through Go's `x/crypto` package.
- Request validation using `go-playground/validator`.
- Environment-based configuration for sensitive settings.

### 🌐 REST API Development

- HTTP API development using the Gin Web Framework.
- JSON-based communication between clients and backend services.
- Modular code organization for maintainability.
- Request binding and validation support.

### 🗄️ Database Management

- MySQL 8.0 as the relational database.
- MySQL database connectivity using `go-sql-driver/mysql`.
- SQL migration management using dbmate.
- Version-controlled database schema changes.

### 🐳 Containerized Database

- Docker Compose configuration for MySQL.
- Persistent database storage using Docker volumes.
- Environment variables for database initialization.
- Isolated database environment for local development.

### 🏗️ Modular Go Architecture

- `cmd/` for application entry points.
- `internal/` for private application logic.
- `pkg/` for reusable packages.
- `db/migrations/` for database schema evolution.

## 🛠️ Tech Stack

| Category | Technologies |
|---|---|
| Programming Language | Go 1.26.1 |
| Web Framework | Gin |
| Database | MySQL 8.0 |
| Database Driver | go-sql-driver/mysql |
| Authentication | JSON Web Tokens (JWT) |
| Validation | go-playground/validator |
| Database Migrations | dbmate |
| Containerization | Docker, Docker Compose |
| Configuration | godotenv |
| Utilities | Google UUID, golang.org/x/crypto |

## 🏗️ Architecture Overview

TwitSocial follows a modular Go backend structure, separating HTTP handling, application logic, and database access.

### Conceptual Request Flow

```text
                  Client
                    |
                    v
               HTTP Request
                    |
                    v
              Gin Web Server
                    |
                    v
             Routing / Middleware
                    |
                    v
               API Handlers
                    |
                    v
             Application Logic
                    |
                    v
              Database Access
                    |
                    v
                 MySQL
```

This diagram illustrates the intended separation of responsibilities rather than an exhaustive mapping of every implementation detail.

## 📁 Project Structure

```text
TwitSocial/
│
├── cmd/                    # Application entry points
│
├── internal/               # Internal application components
│
├── pkg/                    # Reusable packages
│
├── db/
│   └── migrations/         # SQL database migrations
│
├── .junie/
│   └── memory/             # Development metadata
│
├── docker-compose.yml      # MySQL Docker configuration
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── package.json            # dbmate tooling dependency
├── package-lock.json
└── .gitignore
```

## ⚙️ Getting Started

### Prerequisites

Ensure the following tools are installed:

- Go 1.26.1 or a compatible version
- Docker and Docker Compose
- Node.js and npm for the repository's dbmate tooling
- Git

### 1. Clone the Repository

```bash
git clone https://github.com/Prasanna-cpu/TwitSocial.git

cd TwitSocial
```

### 2. Configure Environment Variables

Create a `.env` file in the project root:

```env
# MySQL Docker Configuration
MYSQL_ROOT_PASSWORD=change_me_root
MYSQL_DATABASE=twitsocial
MYSQL_USER=twitsocial_user
MYSQL_PASSWORD=change_me_user

# Database Connection for Local Development
DATABASE_URL=mysql://twitsocial_user:change_me_user@127.0.0.1:3307/twitsocial
```

The MySQL environment variables above correspond to the Docker Compose configuration. The `DATABASE_URL` is an example connection string for migration tooling; application-specific configuration may require additional variables.

Never commit real credentials, secrets, or `.env` files to source control.

### 3. Start MySQL Using Docker Compose

```bash
docker compose up -d
```

The Docker Compose configuration starts a MySQL 8.0 database container named:

```text
twitsocial_db
```

The database is accessible through:

```text
Host: localhost
Port: 3307
Container Port: 3306
```

A named Docker volume, `mysql_data`, persists database files across container restarts.

### 4. Install Go Dependencies

```bash
go mod download
```

Verify the Go module:

```bash
go mod tidy
```

### 5. Run Database Migrations

Install the repository's Node.js development dependencies:

```bash
npm install
```

Apply database migrations using dbmate:

```bash
npx dbmate up
```

Check migration status:

```bash
npx dbmate status
```

Rollback the latest migration:

```bash
npx dbmate down
```

Make sure `DATABASE_URL` is configured correctly before running migrations.

### 6. Run the Application

The repository contains a `cmd/` directory for application entry points.

Run the appropriate Go application package:

```bash
go run ./cmd/<application-directory>
```

Replace `<application-directory>` with the actual executable package under `cmd/`.

### 7. Build the Application

Compile all packages:

```bash
go build ./...
```

Run available Go tests:

```bash
go test ./...
```

Check code formatting:

```bash
gofmt -l .
```

Run static analysis:

```bash
go vet ./...
```

## 🗄️ Database Migrations

TwitSocial uses **dbmate** to manage SQL database migrations.

Migration files are maintained under:

```text
db/migrations/
```

Database migrations allow schema changes to be tracked and applied consistently across development environments.

Common commands:

| Command | Description |
|---|---|
| `npx dbmate up` | Apply pending migrations |
| `npx dbmate down` | Roll back the latest migration |
| `npx dbmate status` | View migration status |
| `npx dbmate new migration_name` | Create a new migration |

## 🔐 Security Considerations

- Keep JWT signing secrets outside source control.
- Validate incoming requests before processing.
- Store user passwords using a suitable password-hashing algorithm.
- Use parameterized SQL queries to reduce SQL injection risks.
- Restrict database credentials and permissions.
- Configure HTTPS for production deployments.
- Avoid exposing development database ports publicly.

## 🧪 Development Commands

```bash
# Download dependencies
go mod download

# Compile Go packages
go build ./...

# Run tests
go test ./...

# Static analysis
go vet ./...

# Start database
docker compose up -d

# Stop database
docker compose down

# Apply migrations
npx dbmate up
```

## 🔮 Future Improvements

- Add OpenAPI/Swagger documentation for all REST endpoints.
- Implement automated CI/CD pipelines using GitHub Actions.
- Add integration tests using a dedicated MySQL test database.
- Introduce Redis caching for frequently requested data.
- Add structured logging and request tracing.
- Containerize the Go API alongside MySQL.
- Introduce rate limiting and additional API security controls.

## 👨‍💻 Author

**Prasanna Kumar M**

- GitHub: [Prasanna-cpu](https://github.com/Prasanna-cpu)
- Repository: [TwitSocial](https://github.com/Prasanna-cpu/TwitSocial)

---

⭐ If you find this project useful, consider giving it a star!
