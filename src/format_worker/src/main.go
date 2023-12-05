package main

import (
	"format_worker/src/image_processing"
	"github.com/nats-io/nats.go"
	"log"
	"path/filepath"
	common "shared"
	"shared/config"
	"strings"
)

func main() {

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
		outputPath := createOutputDir(imagePath)
		log.Println("Changing format work on ", imagePath)
		image_processing.Format(imagePath, outputPath)
		err := conn.Publish(workerConfig.Queues.Output, []byte(outputPath))
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

func getFilenameWithExtension(fullPath, extension string) string {
	fileName := filepath.Base(fullPath)
	result := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + extension
	return result
}

func createOutputDir(imagePath string) string {
	newFilename := getFilenameWithExtension(imagePath, ".png")
	outputPath := filepath.Join("../shared_vol/formatted", newFilename)
	return outputPath
}
