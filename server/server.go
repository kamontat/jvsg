package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var mux = http.NewServeMux()

func AddPath(pattern string, handler func(http.ResponseWriter, *http.Request) (string, error)) {
	mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		isDebug := IsDebug(r)
		isDebugStr := strconv.FormatBool(isDebug)
		w.Header().Set("X-Debug", isDebugStr)

		if isDebug {
			fmt.Printf("server: [%s] %s\n", r.Method, r.URL.Path)
			fmt.Printf("server: headers:\n")
			for headerName, headerValue := range r.Header {
				fmt.Printf("\t%s = %s\n", headerName, strings.Join(headerValue, ", "))
			}
			fmt.Printf("server: request length: %d\n", r.ContentLength)
		}

		response, err := handler(w, r)
		if err != nil {
			log.Printf("server: error: %s", err)
			w.Header().Set("X-Error", err.Error())
			w.WriteHeader(http.StatusBadRequest)

			RequestCounterMetrics.WithLabelValues(r.URL.Path, "400", isDebugStr).Inc()
			RequestDurationMetrics.WithLabelValues(r.URL.Path, "400", isDebugStr).Observe(float64(time.Since(start).Milliseconds()))
			return
		}
		if response != "" && isDebug {
			fmt.Printf("server: %s\n", response)
		}

		RequestCounterMetrics.WithLabelValues(r.URL.Path, "200", isDebugStr).Inc()
		RequestDurationMetrics.WithLabelValues(r.URL.Path, "200", isDebugStr).Observe(float64(time.Since(start).Milliseconds()))
	})
}

func AddPathHandler(path string, handler http.Handler) {
	mux.Handle(path, handler)
}

func StartServer(host string, port string) {
	address := fmt.Sprintf("%s:%s", host, port)

	fmt.Printf("server: listening on %s\n", address)
	server := http.Server{
		Addr:    address,
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
