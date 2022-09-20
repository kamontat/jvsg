FROM bash:5-alpine3.15

RUN apk --no-cache add curl

WORKDIR /app

COPY formatter.txt /app/formatter.txt
COPY index.sh /app/index.sh

ENTRYPOINT [ "/app/index.sh" ]
