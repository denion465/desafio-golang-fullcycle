FROM golang:1.19

WORKDIR /go/src/
ENV PATH="/go/bin:${PATH}"

ENTRYPOINT ["/go/src/.docker/entrypoint.sh"]