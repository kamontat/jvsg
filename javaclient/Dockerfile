FROM gradle:jdk18 AS build

WORKDIR /home/gradle

COPY --chown=gradle:gradle ./build.gradle /home/gradle/build.gradle
COPY --chown=gradle:gradle ./settings.gradle /home/gradle/settings.gradle
COPY --chown=gradle:gradle ./src /home/gradle/src


RUN gradle build --no-daemon
RUN unzip /home/gradle/build/distributions/javaclient.zip

FROM amazoncorretto:18

WORKDIR /app
COPY --from=build /home/gradle/javaclient /app/

ENTRYPOINT ["/app/bin/javaclient"]
