FROM golang:1.23-alpine AS builder

ARG GO_SERVER_DIR=.
ARG GO_SERVER_MAIN_PATH=cmd/app/main.go
ARG GO_SERVER_PORT=8000

WORKDIR /app

COPY $GO_SERVER_DIR .
RUN go mod download
RUN go build $GO_SERVER_MAIN_PATH

FROM golang:1.23-alpine

COPY --from=builder /app/main /app
EXPOSE $GO_SERVER_PORT

ENTRYPOINT ["/app"]