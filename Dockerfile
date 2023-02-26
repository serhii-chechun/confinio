FROM golang:1.20-alpine as builder
WORKDIR /

COPY go.mod .
COPY main.go .

COPY confinio confinio
COPY pkg pkg

RUN go test -count=1 -failfast confinio/...
RUN CGO_ENABLED=0 GOOS=linux go build -o /tmp/confinio .

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /tmp/confinio .

ENTRYPOINT [ "./confinio" ]
