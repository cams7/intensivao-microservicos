FROM golang:1.13 AS builder
WORKDIR /go/src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a --ldflags '-extldflags "-static"' -tags "netgo" -installsuffix netgo -o product

FROM scratch
WORKDIR /app
COPY --from=builder /go/src/products.json .
COPY --from=builder /go/src/product .
CMD ["./product"]