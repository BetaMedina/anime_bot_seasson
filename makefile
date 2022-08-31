#run the application and all it's dependencies with docker
dev:
	docker build -t webhook-microsservice -f docker/dev.dockerfile  .
	docker-compose -f docker/docker-compose.yml up

dev-off:
	docker-compose -f docker/docker-compose.yml down
