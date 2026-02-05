up-all:
	docker-compose up --remove-orphans --force-recreate

down-all:
	docker-compose down -v

build-all:
	docker-compose build

rebuild:
	@if [ -z "$(SERVICE)" ]; then \
		echo "Error: SERVICE parameter is required"; \
		echo "Usage: make rebuild SERVICE=service-name"; \
		echo "Available services: $$(docker compose config --services)"; \
		exit 1; \
	fi
	docker compose build $(SERVICE)
	docker compose up --remove-orphans -d $(SERVICE)

.PHONY: ok error rare slow flaky

ok:
	@curl "http://localhost:8084/mode?mode=ok"

error:
	@curl "http://localhost:8084/mode?mode=error"

rare:
	@curl "http://localhost:8084/mode?mode=rare_error"

slow:
	@curl "http://localhost:8084/mode?mode=slow"

flaky:
	@curl "http://localhost:8084/mode?mode=flaky"
