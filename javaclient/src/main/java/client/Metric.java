package client;

import java.io.IOException;
import java.time.Duration;
import java.time.Instant;

import io.prometheus.client.CollectorRegistry;
import io.prometheus.client.Counter;
import io.prometheus.client.Histogram;
import io.prometheus.client.exporter.PushGateway;

public class Metric {
  private CollectorRegistry registry;
  private PushGateway pushGateway;

  private Counter requestTotal;
  private Histogram requestDuration;
  private Histogram requestWithJsonDuration;

  public Metric(String url) {
    this.registry = new CollectorRegistry();
    this.pushGateway = new PushGateway(url);

    this.requestTotal = Counter.build()
        .name("request_total")
        .help("How many job execute")
        .register(registry);
    this.requestDuration = Histogram.build()
        .name("raw_request_ms")
        .help("How long raw request take")
        .buckets(50, 70, 90, 120, 150, 200, 300, 500)
        .register(registry);
    this.requestWithJsonDuration = Histogram.build()
        .name("json_request_ms")
        .help("How long json request take")
        .buckets(50, 70, 90, 120, 150, 200, 300, 500)
        .register(registry);
  }

  public void newRequest() {
    this.requestTotal.inc();
  }

  public void requestDuration(Instant start) {
    Duration duration = Duration.between(start, Instant.now());
    this.requestDuration.observe(duration.toMillis());
  }

  public void jsonParserDuration(Instant start) {
    Duration duration = Duration.between(start, Instant.now());
    this.requestWithJsonDuration.observe(duration.toMillis());
  }

  public void report(String jobName) throws IOException {
    this.pushGateway.push(this.registry, jobName);
  }
}
