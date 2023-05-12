# syntax=docker/dockerfile:1
FROM golang:1.20.3
WORKDIR /app
ADD . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /hs-auth cmd/main.go
EXPOSE 8080
CMD ["/hs-auth"]