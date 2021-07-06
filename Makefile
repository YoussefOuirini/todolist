serve: ## start images
	docker-compose up -d --remove-orphans

stop: ## stop docker images
	docker-compose stop

down: ## stop docker images
	docker-compose down

destroy: ## stop and remove docker images
	docker-compose down -v

restart:
	make destroy && make serve
