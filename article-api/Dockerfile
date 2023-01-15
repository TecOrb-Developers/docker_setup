FROM golang:1.10-alpine3.8

RUN mkdir /go/src/article-api
WORKDIR /go/src/article-api

RUN apk update && apk add git

#RUN go version && go get -u -v golang.org/x/vgo
#RUN vgo install ./...
#Tried using vgo but couldn't make it work

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/context 
RUN go get github.com/gorilla/handlers
RUN go get gopkg.in/mgo.v2
    
COPY  main.go .
COPY  store ./store

RUN go build -o myapp

ENV PORT=8000
EXPOSE 8000

ENTRYPOINT ["./myapp"]
