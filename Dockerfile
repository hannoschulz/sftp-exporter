FROM registry.access.redhat.com/rhel7 as builder

RUN curl -skL https://golang.org/dl/go1.14.15.linux-amd64.tar.gz -o go1.14.15.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.14.15.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin

WORKDIR /
COPY ./go.mod ./
COPY ./go.sum ./
RUN /usr/local/go/bin/go mod download -x
COPY ./ ./
RUN /usr/local/go/bin/go build -o sftp-exporter

FROM registry.access.redhat.com/rhel7

RUN curl -skL https://golang.org/dl/go1.14.15.linux-amd64.tar.gz -o go1.14.15.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.14.15.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin

WORKDIR /
COPY --from=builder /sftp-exporter .

EXPOSE 8080

ENTRYPOINT ["./sftp-exporter"]
