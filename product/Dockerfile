FROM golang:1.16.2-buster AS builder
WORKDIR /app
COPY ./product/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a --ldflags '-extldflags "-static"' -tags "netgo" -installsuffix netgo -o product

FROM scratch
WORKDIR /app
COPY --from=builder /app/products.json .
COPY --from=builder /app/product .
CMD ["./product"]