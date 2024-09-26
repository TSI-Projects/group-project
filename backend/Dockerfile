FROM golang:1.23-alpine as builder

ARG GO_SERVER_DIR=.
ARG GO_SERVER_MAIN_PATH=main.go

WORKDIR /app

COPY $GO_SERVER_DIR .
RUN go mod download
RUN go build $GO_SERVER_MAIN_PATH

FROM golang:1.23-alpine

COPY --from=builder /app/main /app
EXPOSE 8080

ENTRYPOINT ["/app"]