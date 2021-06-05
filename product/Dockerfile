FROM golang:1.16.2-buster

WORKDIR /app
COPY . .
RUN GOOS=linux go build -ldflags="-s -w"
CMD ["./product"]