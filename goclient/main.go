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
		Collector(requestTotal).
		Collector(requestDuration).
		Collector(requestWithJsonDuration)

	s := gocron.NewScheduler(time.UTC)
	s.Every(executionInterval).Do(func() {
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
	})

	s.Every("30s").Do(func() {
		fmt.Println("Pushing metrics to push gateway")
		err := pusher.Push()
		if err != nil {
			fmt.Println("Could not push to Pushgateway:", err)
		}
	})

	s.StartBlocking()
}
