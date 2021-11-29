# Builder
FROM golang:1.17-alpine as builder

WORKDIR /
COPY . .
RUN apk update && apk upgrade && \
    apk --update add git make
RUN CGO_ENABLED=0 go get -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv
RUN make engine
RUN ls

FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app
WORKDIR /app

EXPOSE 5000 40000

COPY --from=builder ./go/bin/dlv /app/dlv
COPY --from=builder ./engine /app

CMD ["./dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/engine"]