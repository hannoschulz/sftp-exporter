FROM registry.access.redhat.com/rhel7

RUN yum install -y yum-utils && \
    yum -y clean all --enablerepo='*' && \
    curl -skL https://golang.org/dl/go1.14.15.linux-amd64.tar.gz -o go1.14.15.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && tar -C /usr/local -xzf go1.14.15.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin

WORKDIR /
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download -x
COPY ./ ./
RUN go build -o sftp-exporter

# FROM registry.access.redhat.com/rhel7

# WORKDIR /
# COPY --from=builder /sftp-exporter .
ENTRYPOINT ["./sftp-exporter"]
