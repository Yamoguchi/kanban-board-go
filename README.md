# Kanban Board

A learning project for practicing Go, HTTP APIs, PostgreSQL, and Docker.

Planned features include registration, login, personal Kanban boards,
columns, and cards. The PostgreSQL environment and a minimal HTTP server
with a health endpoint are implemented. The API does not connect to the
database yet; the remaining backend features are under development.

## Requirements

- Go 1.27.3
- Docker with Docker Compose support
- GNU Make (`make`)

Docker must be running before starting PostgreSQL.
Run all commands below from the project root.

## Environment Setup

Create a local configuration file:

```bash
cp .env.example .env
```

Set the following variables in `.env`:

- `POSTGRES_DB` — database name.
- `POSTGRES_USER` — database user.
- `POSTGRES_PASSWORD` — local database password.

The `.env` file is excluded from Git. The Makefile requires this file for all
targets, including `make help`, and exports its variables to the commands it runs.

### HTTP Configuration

The HTTP server reads configuration from environment variables using `envconfig`.
Defaults apply when a variable is not set.

| Variable | Default | Purpose |
|---|---|---|
| `HTTP_ADDRESS` | `127.0.0.1:8080` | Server listen address. An explicitly empty value is rejected. |
| `HTTP_READ_HEADER_TIMEOUT` | `5s` | Time allowed to read request headers. |
| `HTTP_READ_TIMEOUT` | `10s` | Time allowed to read the entire request, including its body. |
| `HTTP_WRITE_TIMEOUT` | `10s` | Response write deadline; handler execution also consumes this time. |
| `HTTP_IDLE_TIMEOUT` | `60s` | Time allowed to wait for the next request on a keep-alive connection. |

Timeout values use Go duration syntax, such as `5s`, `500ms`, or `1m`.
Malformed durations cause startup to fail.

## Development Commands

Run `make` or `make help` to list the available commands.

| Command          | Description                                                                      |
|------------------|----------------------------------------------------------------------------------|
| `make env-up`    | Start PostgreSQL in the background and wait for a successful healthcheck.        |
| `make env-down`  | Stop and remove the PostgreSQL container, preserving local database files.       |
| `make db-shell`  | Open the PostgreSQL SQL console inside the container.                            |
| `make clear-env` | Stop PostgreSQL and delete all local database files after explicit confirmation. |
| `make start-http-server` | Run the HTTP API locally using settings from `.env`. |

Start the database:

```bash
make env-up
```

Open the SQL console:

```bash
make db-shell
```

Stop and remove the database container:

```bash
make env-down
```

## Running the HTTP API

Start the server in the foreground:

```bash
make start-http-server
```

With the default address, check the health endpoint from another terminal:

```bash
curl -i http://127.0.0.1:8080/healthz
```

Expected response: HTTP `200 OK`, `Content-Type: application/json`, and
`{"status":"ok"}`. This endpoint checks the HTTP server only, not database
connectivity. PostgreSQL does not need to be running for this endpoint.
Use the configured address in the URL if `HTTP_ADDRESS` differs from the default.

To run the whole Go package directly with an explicit address:

```bash
HTTP_ADDRESS=127.0.0.1:8081 go run ./cmd/api
```

Direct `go run` does not load `.env`; it uses the process environment and defaults.
Stop the server with Ctrl+C.

## Database Storage and Reset

Database files are stored in `out/pgdata` through a bind mount.
The `out/` directory is excluded from Git. Removing or recreating the container
does not delete these files.

To reset the local database, run `make clear-env`. Only `y` or `Y` confirms
deletion; pressing Enter or entering any other answer cancels it.
This permanently deletes all database files in `out/pgdata` and leaves
PostgreSQL stopped. Run `make env-up` to initialize a new database.

The database name, user, and password from `.env` are applied when an empty
data directory is first initialized. Editing `.env` does not change credentials
in an existing database.
