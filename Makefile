compose-up: ## Starts a local instance of the service with instances of the event bus and ArangoDB in docker
	@if [ -f docker-compose.yml ]; then\
		docker-compose -p validation -f docker-compose.yml up -d ;\
	else\
		echo "No compose file.";\
	fi

compose-down: ## Shutdowns the local containers
	@docker-compose -p validation -f docker-compose.yml down

run:
	go run ./main.go