ARG GO_VERSION=1.21
FROM golang:${GO_VERSION} AS build

WORKDIR /src

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/server main.go

COPY ./config.yml /bin/config.yml

FROM alpine:latest AS final

COPY --from=build /bin/server /bin/
COPY --from=build /bin/config.yml /bin/config.yml

EXPOSE 8081
ENTRYPOINT [ "/bin/server" ]
