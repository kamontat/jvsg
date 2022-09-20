package client;

public class Executor implements Runnable {
  private Request requester;

  public Executor(Request requester) {
    this.requester = requester;
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
