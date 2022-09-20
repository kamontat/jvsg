package client;

public class Environment {
  public static String get(String key, String def) {
    String value = System.getenv(key);
    if (value == null)
      return def;
    else
      return value;
  }
}
