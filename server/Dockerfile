FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY deploy .

RUN go build -o server main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/server build/server

EXPOSE 8080

CMD [". /server"]