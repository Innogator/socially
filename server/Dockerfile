FROM golang:1.13.7-alpine
RUN apk add git

COPY . /go/src/socially/server

WORKDIR /go/src/app
COPY ./main.go .
# RUN GO GET 
# 
# 	"github.com/gorilla/mux"
# 	"go.mongodb.org/mongo-driver/bson"
# 	"go.mongodb.org/mongo-driver/bson/primitive"
# 	"go.mongodb.org/mongo-driver/mongo"
# 	"go.mongodb.org/mongo-driver/mongo/options"

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
