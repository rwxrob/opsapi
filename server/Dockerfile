FROM golang:1.18 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .
COPY internal/api ./internal/api
COPY main.go go.mod go.sum go.work ./
ENV CGO_ENABLED=0
RUN go build -o openapi

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]

