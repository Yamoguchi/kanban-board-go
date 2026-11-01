.PHONY: help env-up env-down clear-env db-shell

help:
	@printf '\nKanban Board | Development commands\n\n'
	@printf 'Usage: make <target>\n\n'
	@printf 'Environment\n'
	@printf '  %-14s %s\n' \
		'env-up' 'Start PostgreSQL and wait until healthy' \
		'env-down' 'Stop and remove the PostgreSQL container'
	@printf '\nDatabase\n'
	@printf '  %-14s %s\n' \
		'db-shell' 'Open psql in the PostgreSQL container' \
		'clear-env' 'Delete all PostgreSQL data after confirmation'
	@printf '\nHelp\n'
	@printf '  %-14s %s\n\n' 'help' 'Show available targets'

env-up:
	@docker compose up -d postgres --wait

env-down:
	@docker compose down postgres

clear-env:
	@printf 'Delete all PostgreSQL data? [y/N] '; \
	read -r answer || exit 1; \
	case "$$answer" in \
		Y|y) docker compose stop postgres && rm -rf -- ./out/pgdata ;; \
		*) echo "Cancelled" ;; \
	esac

db-shell:
	@docker compose exec postgres sh -c 'psql -U "$$POSTGRES_USER" -d "$$POSTGRES_DB"'

start-http-server:
	@go run cmd/api/main.go
