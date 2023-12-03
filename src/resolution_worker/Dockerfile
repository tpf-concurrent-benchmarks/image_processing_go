FROM golang:1.21.4

WORKDIR /app

COPY ./resolution_worker/go.mod ./resolution_worker/go.sum ./resolution_worker/
COPY ./common ./common
RUN cd resolution_worker; go mod download

COPY ./resolution_worker/src/ ./resolution_worker/src/

RUN cd resolution_worker/src; CGO_ENABLED=0 GOOS=linux go build -o /app/resolution_worker/resolution-worker

CMD cd /app/resolution_worker/; ./resolution-worker