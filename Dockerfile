FROM golang:1.21-alpine as builder

USER root

WORKDIR /go/src/app

ENV GOEXPERIMENT arenas

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /go/bin/app cmd/main.go

FROM alpine:latest

ENV GIN_MODE release

RUN adduser -D appuser
USER appuser

COPY --from=builder /go/bin/app /go/bin/

CMD ["sh", "-c", "/go/bin/app"]
