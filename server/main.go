package main

import "github.com/prometheus/client_golang/prometheus/promhttp"

const defaultHost = "0.0.0.0"
const defaultPort = "3333"

func main() {
	var host = GetEnv("HOST", defaultHost)
	var port = GetEnv("PORT", defaultPort)

	AddPath("/ping", PingBody)
	AddPath("/mirror", MirrorBody)
	AddPath("/json", GetJsonBody)
	AddPath("/json/cached", GetCacheJsonBody)

	AddPathHandler("/metrics", promhttp.Handler())
	StartServer(host, port)
}
