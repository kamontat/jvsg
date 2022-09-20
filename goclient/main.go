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
	var pushGatewayUrl = GetEnv("PUSH_GATEWAY_URL", defaultPushGateway)
	// "ns", "us" (or "µs"), "ms", "s", "m", "h"
	var executionInterval = GetEnv("EXECUTION_INTERVAL", defaultInterval)
	var path = GetEnv("SERVER_PATH", "/json")
	var debug = GetEnv("SERVER_DEBUG", "false")
	var isParse = GetEnv("PARSE", "true")
	var jobName = GetEnv("JOB_NAME", "goclient")

	pusher := push.New(pushGatewayUrl, jobName)
	pusher.
		Collector(requestCount).
		Collector(requestDuration).
		Collector(requestWithJsonDuration)

	fmt.Printf("Start send request to %s:%s\n", SERVER_HOST, SERVER_PORT)
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(executionInterval).Do(func() {
		NewRequestMetric()
		if isParse == "true" {
			_, err := RequestWithParse(path, debug)
			if err != nil {
				fmt.Println("Could not request to server:", err)
			}
		} else {
			_, err := Request(path, debug)
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
