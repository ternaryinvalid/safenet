ARG GO_VERSION=1.21
FROM golang:${GO_VERSION} AS build

WORKDIR /src

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /bin/client main.go

COPY ./config.yml /bin/config.yml
COPY ./account.json /bin/account.json

FROM alpine:latest AS final

COPY --from=build /bin/client /bin/
COPY --from=build /bin/config.yml /bin/config.yml
COPY --from=build /bin/account.json /bin/account.json

EXPOSE 8000

ENTRYPOINT [ "/bin/client" ]
