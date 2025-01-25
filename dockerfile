FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o MoliDB .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/MoliDB .

EXPOSE 17233

CMD ["./MoliDB"]