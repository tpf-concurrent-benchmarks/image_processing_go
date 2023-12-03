FROM golang:1.21.4

WORKDIR /app

COPY ./size_worker/go.mod ./size_worker/go.sum ./size_worker/
COPY ./common ./common
RUN cd size_worker; go mod download

COPY ./size_worker/src/ ./size_worker/src/

RUN cd size_worker/src; CGO_ENABLED=0 GOOS=linux go build -o /app/size_worker/size-worker

CMD cd /app/size_worker/; ./size-worker