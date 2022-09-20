package client;

public class Executor implements Runnable {
  private Request requester;

  public Executor() {
    this.requester = new Request();
  }

  @Override
  public void run() {
    try {
      this.requester.start();
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
