FROM golang:1.17.5-alpine3.15 AS builder

WORKDIR /api

COPY ./ ./

RUN go mod download

RUN go build ./cmd/api/main.go

FROM alpine:3.15 as production
COPY --from=builder /api/main .
EXPOSE 8080
CMD ["./main"]