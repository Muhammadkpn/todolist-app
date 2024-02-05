# Golang Clean Architecture Boilerplate for Bank Sinarmas

This Golang boilerplate is designed to kickstart your project development using a clean architecture approach. It separates concerns into distinct layers such as handler, usecase, and repository, promoting maintainability and testability.

## Requirements

- [Go](https://golang.org/dl/) (Minimum version: 1.13)
- [Docker](https://www.docker.com/) (Optional but recommended for database and elk setup)
- [Docker Compose](https://docs.docker.com/compose/) (Optional but recommended for local development with Docker)

## Getting Started

1. Clone this repository:

```bash
git clone https://gitlab.banksinarmas.com/go/template.git
cd template
```

2. Install Depedencies

```go
go mod download
```

3. Build Package

```go
make build-http
```

4. Run Package

```bash
./bin/http
```

The application should now be running at http://localhost:3031.

## Project Structure

The project structure follows the clean architecture principles:

```
.
├── bin/
│   └── http                                    # binary file for http service
├── cmd/
│   ├── http                                    # http service commands
│   └── grpc                                    # grpc service commands, not ready yet
├── docker                                      # docker cfg
├── files                                       # files
├── internal/
│   ├── inbound/
│   │   ├── http/                               # http inbound
│   │   │   └── task/
│   │   │       ├── add.go                      # crud file
│   │   │       ├── contract.yaml               # contract for openapi generator
│   │   │       ├── controller.go               # init controller file
│   │   │       └── openapi_server.gen.go       # server side openapi result
│   │   └── model/                              # generated model by open api
│   │       └── task.gen.go
│   ├── repository/                             # repository folder for external call
│   │   ├── db/                                 # db folder
│   │   │   ├── model/                          # model for db
│   │   │   │   └── task.go
│   │   │   └── task/                           # task repository
│   │   │       ├── impl/
│   │   │       │   ├── create.go               # crud file
│   │   │       │   └── init.go                 # init task
│   │   │       └── repository.go               # interface task
│   │   └── redis                               # redis folder
│   ├── usecase/                                # usecase folder
│   │   ├── model/                              # model folder for usecase
│   │   │   └── task.go
│   │   └── task/                               # task usecase
│   │       ├── impl/                           # task implementation
│   │       │   ├── create.go
│   │       │   └── init.go
│   │       └── usecase.go                      # task interface
│   └── util/
│       └── span.go                             # span module
├── pkg/
│   ├── config/                                 # config folder
│   │   ├── config.go                           # init config file
│   │   ├── app.go                              # app struct
│   │   └── database.go                         # database struct
│   ├── constant                                # constant folder
│   ├── di                                      # depedency injection file
│   ├── enums
│   └── resources
├── scripts/
│   └── http.sh
├── tmp
└── variables/
    ├── development.yaml
    └── production.yaml
```

## List of Main Framework and Library

1. [Sinarmas SDK](https://gitlab.banksinarmas.com/go/sdk) - Sinarmas Golang SDK(contains logger, tracer, mandatory)
2. [Echo](https://github.com/labstack/echo) - HTTP web framework
3. [Gorm](https://gorm.io/docs/index.html) - Object Relational Mapping (ORM) library
4. [Apm](https://github.com/elastic/apm-agent-go) - Observability framework
