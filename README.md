# Kanban Board

A learning project for practicing Go, HTTP APIs, PostgreSQL, and Docker.

Planned features include registration, login, personal Kanban boards,
columns, and cards. The PostgreSQL environment is configured;
the backend is under development.

## Requirements

- Go 1.27.3
- Docker with Docker Compose support

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
