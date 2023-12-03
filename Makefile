N_WORKERS=2

init:
	docker swarm init

build:
	docker rmi image_processing_go_format_worker -f
	docker rmi image_processing_go_resolution_worker -f
	docker rmi image_processing_go_size_worker -f
	docker rmi image_processing_go_manager -f
	docker build -t image_processing_go_format_worker ./src/ -f ./src/format_worker/Dockerfile
	docker build -t image_processing_go_resolution_worker ./src/ -f ./src/resolution_worker/Dockerfile
	docker build -t image_processing_go_size_worker ./src/ -f ./src/size_worker/Dockerfile
	docker build -t image_processing_go_manager ./src/ -f ./src/manager/Dockerfile

setup: init build

remove:
	docker stack rm ip_go

docker_build:
	docker compose -f=docker-compose-deploy-local.yml build

create_directories:
	mkdir -p graphite
	mkdir -p shared_vol
	mkdir -p shared_vol/input
	mkdir -p shared_vol/resized
	rm -f shared_vol/resized/*
	mkdir -p shared_vol/formatted
	rm -f shared_vol/formatted/*
	mkdir -p shared_vol/cropped
	rm -f shared_vol/cropped/*

deploy: docker_build create_directories
	N_WORKERS=${N_WORKERS} docker compose -f=docker-compose-deploy-local.yml up

deploy_remote: create_directories
	N_WORKERS=${N_WORKERS} docker stack deploy -c docker-compose-deploy.yml gs_go

down_graphite:
	if docker stack ls | grep -q graphite; then \
		docker stack rm graphite; \
		docker stack rm grafana; \
		docker stack rm cadvisor; \
	fi
.PHONY: down_graphite

format:
	cd ./src/common && go fmt ./...
	cd ./src/manager && go fmt ./...
	cd ./src/format_worker && go fmt ./...
	cd ./src/resolution_worker && go fmt ./...
	cd ./src/size_worker && go fmt ./...
