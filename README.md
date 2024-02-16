# Golang Clean Architecture Boilerplate for Bank Sinarmas

This Golang boilerplate is designed to kickstart your project development using a clean architecture approach. It separates concerns into distinct layers such as handler, usecase, and repository, promoting maintainability and testability.

## Requirements

- [Go](https://golang.org/dl/) (Minimum version: 1.20)
- [Docker](https://www.docker.com/) (Optional but recommended for database and elk setup)
- [Mockery](https://vektra.github.io/mockery/latest/installation) for generate mock
  ```
  go install github.com/vektra/mockery/v2@v2.41.0
  ```
- [OAPICodegen](https://github.com/deepmap/oapi-codegen) for generate openAPI contract into controller
  ```
  go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
  ```

## Key Features

- OpenAPI based contract for consistency, mock, and auto generate controller with [stoplight.io](https://stoplight.io/)
- Dependency Injection with [Uber/Dig](https://github.com/uber-go/dig)
- ORM and Database Transactions using [Gorm](https://gorm.io/docs/index.html)
- Tracing via [APM](https://github.com/elastic/apm-agent-go)
- Logging via [Uber/Zap](https://github.com/uber-go/zap)
- Mocking via [Mockery](https://vektra.github.io/mockery/latest/installation)

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

## How to Use

**1. Define your API Contract:**

Create an OpenAPI specification file named `contract.yaml` within the `./internal/inbound/http/{{your_module}}` directory. Ensure this file aligns with stakeholder expectations and represents the desired API behaviors. You can use tools like Stoplight.io to create and collaborate on this file (creating a free account is required). Guideline to make contract in [here](https://docs.google.com/document/d/1aQeGRRN2Z9am_MhHYKd7r6JdTiy004UbbvfbH_TR7iw/edit?usp=sharing).

**2. Generate Server Code:**

After placing the `contract.yaml` file, run the generator script with `./scripts/http.sh`. This will generate an `openapi_server.gen.go` file in the same directory, containing server-side code based on your API contract.

**3. Implement your Controller:**

Create a `controller.go` file within the same module folder. Define a `Controller` struct that injects any necessary dependencies (indicated by `{{add dependencies here}}`). Additionally, ensure it implements the `StrictServerInterface` interface found in the generated `openapi_server.gen.go` file.

```go
type Controller struct {
        dig.In
        {{add dependencies here}}
    }

    var _ StrictServerInterface = (*Controller)(nil)
```

**4. Implement Controller Functions:**

Inside the generated `openapi_server.gen.go` file, implement the actual logic for each controller function defined in the `StrictServerInterface`. These functions correspond to the endpoints specified in your `contract.yaml` file.

**5. Add Use Cases and Repositories:**

Create separate folders for use cases and repositories as needed to encapsulate business logic and data access concerns. Implement the required functionalities within these folders to complete your application.

**Additional Notes:**

- Replace `{{your_module}}` with the actual name of your module directory.
- Refer to the existing code and comments within the boilerplate for further guidance.
- Consider including instructions on setting up dependencies and running the application.

By following these steps and customizing the provided structure, you can quickly create Go applications with well-defined APIs and organized architecture.
