
run_worker_local:
	cd ./src/worker && ENV=local go run ./src

run_manager_local:
	cd ./src/manager && ENV=local go run ./src

format:
	cd ./src/common && go fmt ./...
	cd ./src/manager && go fmt ./...
	cd ./src/format_worker && go fmt ./...
	cd ./src/resolution_worker && go fmt ./...
	cd ./src/size_worker && go fmt ./...
