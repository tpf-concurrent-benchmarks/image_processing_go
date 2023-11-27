
run_worker_local:
	cd ./src/worker && ENV=local go run ./src

run_manager_local:
	cd ./src/manager && ENV=local go run ./src

format:
	go fmt ./...
