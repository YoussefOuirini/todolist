postgres: ## start infra for the dev env
	docker-compose up -d postgres

migrate: postgres ## run db migrations
	##Giving database some time to start with sleep
	sleep 5
	(go run tools/migrate/main.go up)

migrate-down: postgres
	##Giving database some time to start with sleep
	sleep 5
	(go run tools/migrate/main.go down)

serve: migrate ## start images
	docker-compose up -d --remove-orphans

stop: ## stop docker images
	docker-compose stop

down: ## stop docker images
	docker-compose down

destroy: ## stop and remove docker images
	docker-compose down -v

restart:
	make destroy && make serve
