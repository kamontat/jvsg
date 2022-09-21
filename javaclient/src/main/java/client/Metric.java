package client;

import java.io.IOException;
import java.net.URI;
import java.time.Duration;
import java.time.Instant;

import io.prometheus.client.CollectorRegistry;
import io.prometheus.client.Counter;
import io.prometheus.client.Gauge;
import io.prometheus.client.Histogram;
import io.prometheus.client.exporter.PushGateway;

public class Metric {
  private CollectorRegistry registry;
  private PushGateway pushGateway;

  private Counter requestCount;
  private Gauge requestDuration;
  private Histogram requestDurationBucket;
  private Gauge requestWithJsonDuration;
  private Histogram requestWithJsonDurationBucket;

  public Metric(String name, String link) {
    URI uri = URI.create(link);
    this.registry = new CollectorRegistry();

    String url = String.format("%s:%d", uri.getHost(), uri.getPort());
    this.pushGateway = new PushGateway(url);

    this.requestCount = Counter.build()
        .name(name + "_request_count")
        .help("How many job execute")
        .register(registry);
    this.requestDuration = Gauge.build()
        .name(name + "_raw_request_ms")
        .help("How long raw request take")
        .register(registry);
    this.requestDurationBucket = Histogram.build()
        .name(name + "_raw_request_bucket_ms")
        .help("How long raw request take")
        .buckets(100, 300, 500, 700, 900, 1200, 1500)
        .register(registry);
    this.requestDuration = Gauge.build()
        .name(name + "_json_request_ms")
        .help("How long raw request take")
        .register(registry);
    this.requestWithJsonDurationBucket = Histogram.build()
        .name(name + "_json_request_bucket_ms")
        .help("How long json request take")
        .buckets(100, 300, 500, 700, 900, 1200, 1500)
        .register(registry);
  }

  public void newRequest() {
    this.requestCount.inc();
  }

  public void requestDuration(Instant start) {
    double duration = Duration.between(start, Instant.now()).toMillis();
    this.requestDuration.set(duration);
    this.requestDurationBucket.observe(duration);
  }

  public void jsonParserDuration(Instant start) {
    double duration = Duration.between(start, Instant.now()).toMillis();
    this.requestWithJsonDuration.set(duration);
    this.requestWithJsonDurationBucket.observe(duration);
  }

  public void report(String jobName) throws IOException {
    this.pushGateway.push(this.registry, jobName);
  }
}
