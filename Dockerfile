FROM golang:1.11

COPY . /tmp/src

RUN cd /tmp/src && go install && chmod g+rx /go/bin/pseudoconfigmap

USER 1001
CMD ["/go/bin/pseudoconfigmap"]