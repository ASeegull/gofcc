FROM golang

WORKDIR /go/src/github.com/ASeegull/gofcc/testserver
COPY . /go/src/github.com/ASeegull/gofcc/testserver

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure -vendor-only

RUN go install github.com/ASeegull/gofcc/testserver

ENTRYPOINT /go/bin/testserver

EXPOSE 4000