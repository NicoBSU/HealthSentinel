# syntax=docker/dockerfile:1
FROM golang:1.20.3
WORKDIR /app
ADD . /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /hs-therahab cmd/main.go
EXPOSE 8082
CMD ["/hs-therahab"]