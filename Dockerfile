FROM registry.access.redhat.com/rhel7

RUN yum install -y yum-utils && \
    yum -y install hostname rsync && \
    yum -y install net-tools iproute bind-utils && \
    rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm && \
    yum -y install jq xmlstarlet && \
    yum -y install golang-1.14.15-3.fc32 && \
    yum -y remove epel-release && \
    yum -y clean all --enablerepo='*'

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
