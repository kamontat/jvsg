package client;

import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.ScheduledFuture;
import java.util.concurrent.TimeUnit;

public class App {
    private static final ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(1);

    private static final Executor executor = new Executor();

    public static void main(String[] args) throws Exception {
        String value = Environment.get("EXECUTION_INTERVAL_VALUE", "1");
        String unit = Environment.get("EXECUTION_INTERVAL_UNIT", "SECONDS");

        System.out.println("Start send request");
        ScheduledFuture<?> handle = App.scheduler.scheduleAtFixedRate(executor, 0, Integer.parseInt(value),
                TimeUnit.valueOf(unit));

        Runnable canceller = () -> handle.cancel(false);
        scheduler.schedule(canceller, 24, TimeUnit.HOURS);
    }
}
