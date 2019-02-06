#hiveon-api image
FROM registry.tor.ph/go/hiveon-api:base as build-deps
RUN mkdir -p /app/src/hiveon-api && \
    mkdir -p /app/swaggerui
COPY . /app/src/hiveon-api
WORKDIR /app/src/hiveon-api
RUN mv conf/config.dev.yaml conf/config.yaml && \
    go get ./... && \
    make swagger && \
    go build ./... 

FROM golang:alpine3.8
RUN mkdir -p /app/conf
WORKDIR /app
COPY --from=build-deps /app/bin/hiveon-api /app
COPY --from=build-deps /app/src/hiveon-api/conf/. /app/conf
COPY --from=build-deps /app/src/hiveon-api/swaggerui /app/swaggerui
ENV build-number=${CI_PIPELINE_ID:-latest}
EXPOSE 8080 8090
CMD ["./hiveon-api"]
