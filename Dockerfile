FROM golang:alpine

EXPOSE 8000

ENTRYPOINT [ "./2do" ]

RUN mkdir /2do

WORKDIR /2do

COPY . /2do

RUN mkdir db

RUN go get -v -d

RUN go build .