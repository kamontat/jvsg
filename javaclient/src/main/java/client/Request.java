package client;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpRequest.BodyPublishers;
import java.net.http.HttpResponse.BodyHandlers;
import java.time.Instant;

public class Request {
  private HttpClient client;
  private Metric metrics;

  private String host;
  private String port;
  private String path;
  private String debug;
  private String parse;
  private String jobName;

  public Request() {
    this.client = HttpClient.newHttpClient();

    this.host = Environment.get("SERVER_HOST", "localhost");
    this.port = Environment.get("SERVER_PORT", "3333");
    this.path = Environment.get("SERVER_PATH", "/json");
    this.debug = Environment.get("SERVER_DEBUG", "false");
    this.parse = Environment.get("PARSE", "false");
    this.jobName = Environment.get("JOB_NAME", "javaclient");

    String pushGatewayUrl = Environment.get("PUSH_GATEWAY_URL", "http://localhost:9091");
    this.metrics = new Metric(this.jobName, pushGatewayUrl);
  }

  public String getServer() {
    return String.format("%s:%s", this.host, this.port);
  }

  public void start() throws Exception {
    this.metrics.newRequest();

    var url = "http://" + host + ":" + port + path + "?debug=" + debug;
    var request = HttpRequest.newBuilder(URI.create(url))
        .POST(BodyPublishers.noBody())
        .build();

    if (parse == "true") {
      var start = Instant.now();
      var response = client.send(request, new JsonBodyHandler());
      this.metrics.requestWithJsonDuration(start);
      if (response.statusCode() != 200) {
        throw new Exception("Unexpected status code: " + response.statusCode());
      }
    } else {
      var start = Instant.now();
      var response = client.send(request, BodyHandlers.ofString());
      this.metrics.requestDuration(start);
      if (response.statusCode() != 200) {
        throw new Exception("Unexpected status code: " + response.statusCode());
      }
    }

    this.metrics.report(this.jobName);
  }
}
