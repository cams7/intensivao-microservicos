FROM golang:1.16.2-buster AS builder
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a --ldflags '-extldflags "-static"' -tags "netgo" -installsuffix netgo -o catalog

FROM scratch
WORKDIR /app
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/catalog .
CMD ["./catalog"]