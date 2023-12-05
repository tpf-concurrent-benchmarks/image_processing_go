package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"manager/src/utils"
	common "shared"
	"shared/config"
	"time"
	"github.com/cactus/go-statsd-client/v5/statsd"
)

func main() {
	managerConfig := config.GetConfig()
	connAddress := config.CreateConnectionAddress(managerConfig.Host, managerConfig.Port)

	metricsAddr := config.CreateMetricAddress(managerConfig.Metrics.Host, managerConfig.Metrics.Port)
	statsdClient := CreateStatsClient(metricsAddr, utils.GetNodeID())

	startTime := time.Now()

	natsConnection, err := nats.Connect(connAddress)
	if err != nil {
		log.Fatalf("Unable to connect to NATS: %v", err)
	}
	imagesFiles := utils.GetImagesInDirectory("../shared_vol/input/")
	shouldStop := make(chan bool)

	// For sync purposes
	time.Sleep(5 * time.Second)

	subscribeForResults(natsConnection, managerConfig.Queues.Input, len(imagesFiles), shouldStop)

	sendWork(natsConnection, managerConfig.Queues.Output, imagesFiles)

	if <-shouldStop {
		log.Println("All results received, sending end message to all workers")
		sendEndMessage(natsConnection, common.EndWorkMessage, common.EndWorkQueue)

		endTime := time.Now()
		elapseTime := endTime.Sub(startTime).Milliseconds()
		err = statsdClient.Gauge("completion_time", elapseTime, 1.0)
		if err != nil {
			log.Fatalf("Error sending metric to statsd: %s", err)
		}

		close(shouldStop)
		natsConnection.Close()
	}

}

func sendEndMessage(connection *nats.Conn, message string, endQueue string) {
	publishMessage(connection, message, endQueue)
}

func publishMessage(connection *nats.Conn, message string, s string) {
	err := connection.Publish(s, []byte(message))
	if err != nil {
		log.Fatalf("Unable to publish message: %v", err)
	}
}

func sendWork(connection *nats.Conn, outputQueue string, files []string) {
	for _, file := range files {
		publishMessage(connection, file, outputQueue)
	}
}

func subscribeForResults(connection *nats.Conn, inputQueue string, workAmount int, stop chan bool) {
	resultsReceived := 0
	log.Println("Subscribing for results, work sent: ", workAmount)
	_, _ = connection.Subscribe(inputQueue, func(msg *nats.Msg) {
		message := string(msg.Data)
		if message != common.JobDoneMessage {
			log.Println("Ignoring message from worker: ", message)
		}
		resultsReceived++
		if resultsReceived == workAmount {
			stop <- true
		}
	})
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