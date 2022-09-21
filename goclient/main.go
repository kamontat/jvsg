package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/prometheus/client_golang/prometheus/push"
)

const defaultPushGateway = "http://localhost:9091"
const defaultInterval = "1s"

func main() {
	pusher := push.New(pushGatewayUrl, jobName)
	pusher.
		Collector(requestCount).
		Collector(requestDuration).
		Collector(requestDurationBucket).
		Collector(requestWithJsonDuration).
		Collector(requestWithJsonDurationBucket)

	fmt.Printf("Start send request to %s:%s\n", serverHost, serverPort)
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(executionInterval).Do(func() {
		NewRequestMetric()
		if isParse == "true" {
			_, err := RequestWithParse(path, debug)
			if err != nil {
				fmt.Println("Could not request to server:", err)
			}
		} else {
			_, err := RequestWithoutParse(path, debug)
			if err != nil {
				fmt.Println("Could not request to server:", err)
			}
		}

		err := pusher.Push()
		if err != nil {
			fmt.Println("Could not push to Pushgateway:", err)
		}
	})

	if err != nil {
		fmt.Println("Could not create job:", err)
	}

	s.StartBlocking()
	fmt.Println("End send request")
}
