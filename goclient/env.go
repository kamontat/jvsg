package main

var (
	pushGatewayUrl    = GetEnv("PUSH_GATEWAY_URL", defaultPushGateway)
	executionInterval = GetEnv("EXECUTION_INTERVAL", defaultInterval)
	path              = GetEnv("SERVER_PATH", "/json")
	debug             = GetEnv("SERVER_DEBUG", "false")
	isParse           = GetEnv("PARSE", "false")
	jobName           = GetEnv("JOB_NAME", "goclient")
	serverHost        = GetEnv("SERVER_HOST", "localhost")
	serverPort        = GetEnv("SERVER_PORT", "3333")
)
