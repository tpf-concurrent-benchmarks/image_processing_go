package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"path/filepath"
	common "shared"
	"shared/config"
	"size_worker/src/image_processing"
	"github.com/cactus/go-statsd-client/v5/statsd"
	"size_worker/src/utils"
	"time"
)

func main() {
	workerConfig := config.GetConfig()
	connString := config.CreateConnectionAddress(workerConfig.Host, workerConfig.Port)
	natsConn, err := nats.Connect(connString)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %s", err)
	}

	metricsAddr := config.CreateMetricAddress(workerConfig.Metrics.Host, workerConfig.Metrics.Port)
	statsdClient := CreateStatsClient(metricsAddr, utils.GetNodeID())

	shouldStop := make(chan bool)

	waitForEnd(natsConn, common.EndWorkQueue, shouldStop)

	subscribeForWork(natsConn, workerConfig, statsdClient)

	if <-shouldStop {
		natsConn.Close()
		close(shouldStop)
	}
}

func subscribeForWork(conn *nats.Conn, workerConfig config.Config, statsdClient statsd.Statter) {
	_, err := conn.QueueSubscribe(workerConfig.Queues.Input, "workers_group", func(msg *nats.Msg) {
		imagePath := string(msg.Data)
		newImagePath := createOutputDir(imagePath)

		startTime := time.Now()

		image_processing.CropCentered(imagePath, newImagePath, workerConfig.Worker.TargetWidth, workerConfig.Worker.TargetHeight)

		endTime := time.Now()
		elapseTime := endTime.Sub(startTime).Milliseconds()
		err := statsdClient.Timing("work_time", elapseTime, 1.0)
		if err != nil {
			log.Fatalf("Error sending metric to statsd: %s", err)
		}

		err = statsdClient.Inc("results_produced", 1, 1.0)
		if err != nil {
			log.Fatalf("Error sending metric to statsd: %s", err)
		}

		err = conn.Publish(workerConfig.Queues.Output, []byte(common.JobDoneMessage))
		if err != nil {
			log.Fatalf("Error publishing to queue: %s", err)
		}
	})
	if err != nil {
		log.Fatalf("Error subscribing to queue: %s", err)
	}
}

func waitForEnd(conn *nats.Conn, endQueue string, stop chan bool) {
	_, err := conn.Subscribe(endQueue, func(msg *nats.Msg) {
		message := string(msg.Data)
		if message == common.EndWorkMessage {
			log.Println("Received end message")
			stop <- true
		}
	})
	if err != nil {
		log.Fatalf("Error subscribing to end queue: %s", err)
	}
}

func createOutputDir(imagePath string) string {
	filename := filepath.Base(imagePath)
	outputPath := filepath.Join("../shared_vol/cropped", filename)
	return outputPath
}

func CreateStatsClient(metricsAddr, prefix string) statsd.Statter {
	clientConfig := &statsd.ClientConfig{
		Address: metricsAddr,
		Prefix:  prefix,
	}

	statsdClient, err := statsd.NewClientWithConfig(clientConfig)
	if err != nil {
		log.Fatalf("Error creating statsd client: %s", err)
	}
	return statsdClient
}