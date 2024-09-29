FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV PORT 8080

WORKDIR /app/cmd/app

RUN go build -o /app/bin/braciate .

EXPOSE 8080

CMD ["/app/bin/braciate"]
