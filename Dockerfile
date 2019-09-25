FROM golang

ADD . /go/src/github.com/beegons

WORKDIR /go/src/github.com/beegons

RUN go install github.com/beegons

ENTRYPOINT /go/bin/beegons

EXPOSE 9000