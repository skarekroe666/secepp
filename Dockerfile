FROM golang:1.25.4-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN apk add --no-cache gcc libc-dev
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-s -w" -o secepp .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/secepp .

# RUN go build

CMD [ "./secepp" ]

