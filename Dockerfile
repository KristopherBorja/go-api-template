FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

FROM alpine:latest AS deployment

RUN apk --no-cache add ca-certificates

WORKDIR /go-api-template

COPY --from=builder /app/app .
COPY --from=builder /app/config.json ./config.json

EXPOSE 8080

CMD ["./app"]
