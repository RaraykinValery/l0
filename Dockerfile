# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /orders_service ./cmd/service/

RUN chmod 755 entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]

CMD ["/orders_service"]
