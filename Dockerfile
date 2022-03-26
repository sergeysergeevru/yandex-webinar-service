FROM golang:1.18 as builder

ARG VERSION=unknown

WORKDIR /app
COPY . .
RUN GO111MODULE=on GOOS=linux CGO_ENABLED=0 \
    go build -ldflags "-s -w -X main.version=${VERSION}" \
    -o /app/build/cmd cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/build/cmd /bin/cmd

ENTRYPOINT ["cmd"]

