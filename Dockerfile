FROM golang:1.14-buster as builder

WORKDIR /sftp-exporter
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download -x
COPY ./ ./
RUN go build -ldflags "-X main.version=dev" -o out/sftp-exporter main.go

FROM debian:buster

WORKDIR /sftp-exporter
COPY --from=builder /sftp-exporter/bin/sftp-exporter .
ENTRYPOINT ["./sftp-exporter"]
