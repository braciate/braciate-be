FROM golang:alpine AS builder

RUN apk update

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux  go build -o bin/braciate cmd/app/main.go

FROM alpine AS production

RUN apk update && apk add --no-cache ca-certificates make && update-ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/braciate .
COPY --from=builder /app/.prod.env .env
RUN cat .env
COPY --from=builder /app/database/migrations ./database/migrations
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

RUN chmod +x /usr/local/bin/migrate

ENTRYPOINT ["./braciate"]

FROM alpine AS staging

RUN apk update && apk add --no-cache ca-certificates make && update-ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/braciate .
COPY --from=builder /app/.staging.env .env
COPY --from=builder /app/database/migrations ./database/migrations
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

RUN chmod +x /usr/local/bin/migrate

ENTRYPOINT ["./braciate"]
