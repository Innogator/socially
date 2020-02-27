FROM golang:1.13.7

COPY . /go/src/socially/server

RUN go get "github.com/gorilla/mux"
RUN go get "go.mongodb.org/mongo-driver/bson"
RUN go get "go.mongodb.org/mongo-driver/bson/primitive"
RUN go get "go.mongodb.org/mongo-driver/mongo"
RUN go get "go.mongodb.org/mongo-driver/mongo/options"

RUN go install socially/server

ENTRYPOINT /go/bin/server

EXPOSE 8080
