FROM golang:1.22-alpine as builder
ENV CGO_ENABLED=0

WORKDIR /app

COPY . .
RUN go mod download


WORKDIR /app
RUN go build -o /bin/auth ./cmd/auth/main.go

FROM alpine

WORKDIR /bin

COPY --from=builder /bin/auth /bin/auth
COPY .env .env
COPY ./config ./config

CMD ["/bin/auth"]