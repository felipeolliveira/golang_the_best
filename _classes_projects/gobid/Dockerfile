#
# Building
#
FROM golang:1.23-alpine AS builder
WORKDIR /app

RUN go install github.com/jackc/tern@latest

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

#
# Running
#
FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache postgresql bash

COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /go/bin/tern /usr/local/bin/tern
COPY --from=builder /app/internal/store/pgstore/migrations /app/internal/store/pgstore/migrations

RUN mkdir -p /app/scripts
COPY scripts/init.sh /app/scripts/init.sh
RUN chmod +x /app/scripts/init.sh

ARG GOBID_APP_PORT=8080
EXPOSE ${GOBID_APP_PORT}

CMD ["/app/scripts/init.sh"]
