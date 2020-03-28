FROM golang:1.14-alpine3.11 AS builder

# we need kafka libs in order to run the producer and consumer binaries
RUN apk update && apk add --no-cache librdkafka-dev pkgconf
WORKDIR /project
COPY . .

# FIXME the go build won't work as librdkafka-dev lib is not up to date on Alpine -> we need to get the v1.3.0 lib and Alpine has v1.2.x
#    use confluent repositories to fix this

# Building binaries from builder stage
FROM builder as producer
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o autoreflex-ad-url-producer ./cmd/ad-url-producer

# and the other one with producer image to use the go build cache
FROM producer as consumer
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o autoreflex-ad-url-consumer ./cmd/ad-url-consumer

FROM alpine:3.11
RUN apk update && apk add --no-cache librdkafka
COPY --from=producer /project/autoreflex-ad-url-producer .
COPY --from=consumer /project/autoreflex-ad-url-consumer .
ENTRYPOINT ./${MODE}
