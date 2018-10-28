FROM golang:1.10.3-alpine3.7

COPY . /go/src/github.com/malefaro/technopark-db-forum
WORKDIR /go/src/github.com/malefaro/technopark-db-forum
ENV gopath /go
RUN cd /go/src/github.com/malefaro/technopark-db-forum && go build -o apiserver
EXPOSE 5000
CMD ["/go/src/github.com/malefaro/technopark-db-forum/apiserver"]