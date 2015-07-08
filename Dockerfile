FROM golang:1.4-wheezy

RUN go get github.com/tools/godep
COPY . /go/src/github.com/bountylabs/etcd_groupcache
RUN cd /go/src/github.com/bountylabs/etcd_groupcache/example && godep go build -o /go/bin/main
RUN cp /go/src/github.com/bountylabs/etcd_groupcache/example/cmd.sh /go/bin/cmd.sh
EXPOSE 8080
ENTRYPOINT /go/bin/cmd.sh
