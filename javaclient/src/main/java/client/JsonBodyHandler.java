package client;

import java.io.IOException;
import java.io.InputStream;
import java.net.http.HttpResponse;
import java.net.http.HttpResponse.BodySubscriber;
import java.net.http.HttpResponse.ResponseInfo;
import java.util.Arrays;
import java.util.List;
import java.util.Map;

import com.fasterxml.jackson.core.exc.StreamReadException;
import com.fasterxml.jackson.databind.DatabindException;
import com.fasterxml.jackson.databind.ObjectMapper;

public class JsonBodyHandler implements HttpResponse.BodyHandler<List<Map<?, ?>>> {
  private ObjectMapper mapper;

  JsonBodyHandler() {
    this.mapper = new ObjectMapper();
  }

  @Override
  public BodySubscriber<List<Map<?, ?>>> apply(ResponseInfo responseInfo) {
    HttpResponse.BodySubscriber<InputStream> upstream = HttpResponse.BodySubscribers.ofInputStream();
    return HttpResponse.BodySubscribers.<InputStream, List<Map<?, ?>>>mapping(upstream,
        (InputStream body) -> {
          try {
            List<Map<?, ?>> result = Arrays.asList(mapper.readValue(body, Map[].class));
            return result;
          } catch (StreamReadException e) {
            e.printStackTrace();
          } catch (DatabindException e) {
            e.printStackTrace();
          } catch (IOException e) {
            e.printStackTrace();
          }

          return List.of();
        });
  }
}
