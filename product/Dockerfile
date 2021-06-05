FROM golang:1.13

WORKDIR /go/src/
COPY ./product/ .
RUN GOOS=linux go build -ldflags="-s -w"
CMD ["./product"]