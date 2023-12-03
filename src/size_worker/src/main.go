package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"path/filepath"
	common "shared"
	"shared/config"
	"size_worker/src/image_processing"
)

func main() {
	image_processing.CropCentered("src/resources/rust.png", "src/resources/rust_cropped.png", 100, 100)

	workerConfig := config.GetConfig()
	connString := config.CreateConnectionAddress(workerConfig.Host, workerConfig.Port)
	natsConn, err := nats.Connect(connString)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %s", err)
	}

	//metricsAddr := config.CreateMetricAddress(workerConfig.Metrics.Host, workerConfig.Metrics.Port)
	shouldStop := make(chan bool)

	waitForEnd(natsConn, common.EndWorkQueue, shouldStop)

	subscribeForWork(natsConn, workerConfig)

	if <-shouldStop {
		natsConn.Close()
		close(shouldStop)
	}
}

func subscribeForWork(conn *nats.Conn, workerConfig config.Config) {
	_, err := conn.QueueSubscribe(workerConfig.Queues.Input, "workers_group", func(msg *nats.Msg) {
		imagePath := string(msg.Data)
		// TODO: replace with image cropping
		log.Println("Simulating work on ", imagePath)
		//newImagePath := createOutputDir(imagePath)
		err := conn.Publish(workerConfig.Queues.Output, []byte(common.JobDoneMessage))
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
