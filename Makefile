.PHONY: build run stop clean prune

build:
	@echo "Building Docker images..."
	sudo docker-compose build

run:
	@echo "Starting Docker containers..."
	sudo docker-compose up -d

stop:
	@echo "Stopping running containers..."
	@if [ ! -z "$$(sudo docker ps -q)" ]; then \
		sudo docker stop $$(sudo docker ps -q); \
	else \
		echo "No running containers to stop."; \
	fi

clean:
	@echo "Stopping and removing all containers..."
	@if [ ! -z "$$(sudo docker ps -aq)" ]; then \
		sudo docker stop $$(sudo docker ps -aq); \
		sudo docker rm $$(sudo docker ps -aq); \
	else \
		echo "No containers to clean."; \
	fi

# Remove all unused networks
clean-networks:
	@echo "Removing unused Docker networks..."
	@if [ ! -z "$$(docker network ls -q)" ]; then \
		sudo docker network rm $$(docker network ls -q); \
	else \
		echo "No networks to remove."; \
	fi

# Remove all unused volumes
clean-volumes:
	@echo "Removing all Docker volumes..."
	@if [ ! -z "$$(docker volume ls -q)" ]; then \
		sudo docker volume rm $$(docker volume ls -q); \
	else \
		echo "No volumes to remove."; \
	fi

# Prune Docker system (removes unused images, containers, volumes, etc.)
prune:
	@echo "Pruning Docker system..."
	sudo docker system prune -a --volumes -f

