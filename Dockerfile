FROM golang:1.10
RUN mkdir -p /go/src/github.com/gtongy/demo-echo-app
ADD . /go/src/github.com/gtongy/demo-echo-app
WORKDIR /go/src/github.com/gtongy/demo-echo-app
RUN go get -u github.com/golang/dep/...
RUN dep ensure