FROM golang:1.21.4 as builder
LABEL stage=builder

WORKDIR /app

COPY ./manager/go.mod ./manager/go.sum ./manager/
COPY ./common ./common
RUN cd manager; go mod download

COPY ./manager/src/ ./manager/src/

RUN cd manager/src; CGO_ENABLED=0 GOOS=linux go build -o /app/manager/manager

FROM alpine

WORKDIR /app

COPY --from=builder /app /app

CMD cd /app/manager/; ./manager