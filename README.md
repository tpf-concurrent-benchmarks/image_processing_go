# Image processing in Go

## Running the project locally

The minimum required version of `Go` is 1.21.4 as per requested in the `go.mod` file.
There are four projects in this repository. Either of them can be built with the following commands:

```bash
cd ./src/XX && LOCAL=local go run ./src
```
where `XX` is either `manager`, `format_worker`, `resolution_worker` or `size_worker`.


## Running all services with Docker

```bash
docker compose -f=docker-compose-deploy-local.yml up --build
```

## Number of worker replicas

If you wish to change the number of worker replicas, you can do so by changing the `N_WORKERS` constant in the `Makefile` file.

## Makefile

There is a Makefile in the root directory of the project that can be used to build and run the project

- `make build`: builds manager and worker images.
- `make deploy`: deploys the manager and worker services locally, alongside with Graphite, Grafana and cAdvisor.
- `make deploy_remote`: deploys (with Docker Swarm) the manager and worker services, alongside with Graphite, Grafana and cAdvisor.
- `make remove`: removes all services (stops the swarm).
- `make run_manager_local:` runs the manager locally. Same for `make run_resolution_worker_local`, `make run_size_worker_local` and `make run_format_worker`

## Used libraries

- [Statsd client](https://github.com/cactus/go-statsd-client): used to send metrics to Graphite.
- [NATS client](https://github.com/nats-io/nats.go): used to communicate between manager and worker.
- [Imaging](https://github.com/disintegration/imaging): used to image processing.