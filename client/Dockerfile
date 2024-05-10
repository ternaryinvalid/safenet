FROM golang:alpine as builder
WORKDIR /build

COPY go.mod .

RUN go mod download

COPY deploy .

RUN CGO_ENABLED=0 go build -o /main main.go

FROM scratch

COPY --from=builder main /bin/main

ENTRYPOINT ["/bin/main"]