NAME = mongo4.4

.PHONY = run stop clean

BINARY_NAME=test-app
# run-mongo:
# 	docker run -d -p 8080:27017 --name $(NAME) mongo:4.4

# build-local:
# 	docker build -t "test-app-image" .

# ./bin/test-app:
# 	@echo "Rebuilding..."
# 	@go build -o ./bin/$(BINARY_NAME) .
build:
	@echo "Building..."
	@go build -o ./out/$(BINARY_NAME).out .

run-dev: build
	@echo "Docker Compose Up..."
	@docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d
	@echo "Running..."
	@./out/$(BINARY_NAME).out --env dev

run-prod: build
	@echo "Docker Compose Up..."
	@docker compose -f docker-compose.yml -f docker-compose.prod.yml up -d
	@echo "Running..."
	@./out/$(BINARY_NAME).out --env prod

run: build
	@echo "Docker Compose Up..."
	@docker compose up -d
	@echo "Running..."
	@./out/$(BINARY_NAME).out
	
stop:
	@echo "Docker Compose Down..."
	docker compose down

clean: stop
	@echo "Removing Files..."
	rm -rf ./out
	@echo "Removing Containers..."
	docker compose rm -s

