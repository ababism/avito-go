include .env

# UP

up:
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait

up-all:
	docker-compose -f ./deployments/compose.yaml up -d --build
	docker compose --file ./docker-compose.yml --env-file ./.env up -d --build --wait

up-banner:
	docker-compose up --build banner-svc

up-b:
	docker-compose up --build banner-svc

up-d:
	docker-compose up --build banner-svc -d

# DOWN

down:
	docker-compose down banner-svc

down-c:
	docker-compose down banner-svc

# Observability

up-obs:
	docker-compose -f ./deployments/compose.yaml up -d --build

down-obs:
	docker-compose -f ./deployments/compose.yaml down

# Migrations

migrate-up:
	migrate -path ./services/banner/migrations -database 'postgres://$(BANNER_POSTGRES_USER):$(BANNER_POSTGRES_PASSWORD)@$(BANNER_POSTGRES_HOST_LOCAL):$(BANNER_POSTGRES_PORT_EXTERNAL)/$(BANNER_POSTGRES_NAME)?sslmode=disable' up

migrate-down:
	migrate -path ./services/banner/migrations -database 'postgres://$(BANNER_POSTGRES_USER):$(BANNER_POSTGRES_PASSWORD)@$(BANNER_POSTGRES_HOST_LOCAL):$(BANNER_POSTGRES_PORT_EXTERNAL)/$(BANNER_POSTGRES_NAME)?sslmode=disable' down 1
