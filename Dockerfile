FROM golang:alpine

EXPOSE 8000/tcp

ENTRYPOINT ["todo"]

RUN \
    apk add --update git && \
    rm -rf /var/cache/apk/*

RUN mkdir -p /usr/local/go/src/todo
WORKDIR /usr/local/go/src/todo

COPY . /usr/local/go/src/todo

RUN go get -v -d
RUN go get github.com/GeertJohan/go.rice/rice
RUN rice embed-go
RUN go install -v
