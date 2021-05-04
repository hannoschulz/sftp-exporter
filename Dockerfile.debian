FROM golang:1.14-buster as builder

WORKDIR /
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download -x
COPY ./ ./
RUN go build -o sftp-exporter

FROM debian:buster

WORKDIR /
COPY --from=builder /sftp-exporter .
ENTRYPOINT ["./sftp-exporter"]
