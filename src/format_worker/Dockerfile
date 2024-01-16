FROM golang:1.21.4 as builder
LABEL stage=builder

WORKDIR /app

COPY ./format_worker/go.mod ./format_worker/go.sum ./format_worker/
COPY ./common ./common
RUN cd format_worker; go mod download

COPY ./format_worker/src/ ./format_worker/src/

RUN cd format_worker/src; CGO_ENABLED=0 GOOS=linux go build -o /app/format_worker/format-worker

FROM alpine

WORKDIR /app

COPY --from=builder /app /app

CMD cd /app/format_worker/; ./format-worker