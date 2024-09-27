#Docker Pipeline
FROM alpine:3.17 as builder
WORKDIR /app
COPY . /app
RUN apk add go git && \
    go mod download && \
    go build -o ./bin/music_service

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/bin/music_service ./bin/music_service
EXPOSE 9000
ENTRYPOINT ["./bin/music_service"]
