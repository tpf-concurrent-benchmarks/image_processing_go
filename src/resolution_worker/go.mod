module resolution_worker

go 1.21

require (
	github.com/nats-io/nats.go v1.31.0
	shared v0.0.0-00010101000000-000000000000
)

require (
	github.com/cactus/go-statsd-client/v5 v5.1.0 // indirect
	github.com/disintegration/imaging v1.6.2 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/nats-io/nkeys v0.4.5 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/image v0.14.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)

replace shared => ../common
