FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN docker compose down

RUN docker compose up -d

WORKDIR /app/cmd

RUN go mod tidy

RUN go build -o binary

ENTRYPOINT ["/app/cmd/binary"]