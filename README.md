# Image processing in Go

## Objective

This is a Go implementation of an image processing pipeline under [common specifications](https://github.com/tpf-concurrent-benchmarks/docs/tree/main/image_processing) defined for multiple languages.

The objective of this project is to benchmark the language on a real-world distributed system.

## Deployment

### Requirements

- [Docker >3](https://www.docker.com/) (needs docker swarm)
- [Golang >1.21.4](https://golang.org/)

### Configuration

- **Number of replicas:** `N_WORKERS` constant is defined in the `Makefile` file (this config is built into the containers).
- **Manager config:** in `src/manager/src/resources/config.json` you can define (this config is built into the container):
  - middleware config (nats addres, queues, etc)
  - metrics config (graphite address)

### Commands

#### Startup

- `make setup`: runs both `init` and `build`.
  - `make init`: starts docker swarm.
  - `make build`: builds manager and worker images.
- `template_data`: downloads test image into the input folder

#### Run

- `make deploy`: deploys the manager and worker services locally, alongside with Graphite, Grafana and cAdvisor.
- `make remove`: removes all services (stops the swarm).

> If the manager fails because a service is not ready, increase `SYNC_TIME` in the Makefile and retry.

#### Logs

- `make manager_logs`: shows the logs of the manager service
- `make format_logs`, `make res_logs`, `make size_logs`: shows the logs of the worker services

#### Running the project locally

The minimum required version of `Go` is 1.21.4 as per requested in the `go.mod` file.
There are four projects in this repository. Either of them can be built with the following commands:

```bash
cd ./src/XX && LOCAL=local go run ./src
```

where `XX` is either `manager`, `format_worker`, `resolution_worker` or `size_worker`.

## Libraries

- [Statsd client](https://github.com/cactus/go-statsd-client): used to send metrics to Graphite.
- [NATS client](https://github.com/nats-io/nats.go): used to communicate between manager and worker.
- [Imaging](https://github.com/disintegration/imaging): used to image processing.
