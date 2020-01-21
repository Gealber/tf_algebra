FROM golang:alpine as builder
RUN apk update && apk add --no-cache git \
    && apk add nodejs

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg

ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH 

RUN mkdir -p $GOPATH/src/app 

ADD . $GOPATH/src/app
WORKDIR $GOPATH/src/app 

RUN go build -o api . 

FROM leafney/alpine-mongo:latest


CMD ["/go/src/app/api"]