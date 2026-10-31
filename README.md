# Kanban Board

A learning project for practicing Go, HTTP APIs, PostgreSQL, and Docker.

Planned features include registration, login, personal Kanban boards,
columns, and cards. The PostgreSQL environment is configured;
the backend is under development.

## Requirements

- Go 1.27.3
- Docker with Docker Compose support
- GNU Make (`make`)

Docker must be running before starting the environment.
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

The `.env` file is excluded from Git.

## Development Commands

Run `make` or `make help` to list the available commands.

| Command          | Description                                                                      |
|------------------|----------------------------------------------------------------------------------|
| `make env-up`    | Start PostgreSQL in the background and wait for a successful healthcheck.        |
| `make env-down`  | Stop and remove the PostgreSQL container, preserving local database files.       |
| `make db-shell`  | Open the PostgreSQL SQL console inside the container.                            |
| `make clear-env` | Stop PostgreSQL and delete all local database files after explicit confirmation. |

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
