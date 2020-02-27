FROM golang:1.13.7 as build

COPY . /go/src/socially/server

RUN go get "github.com/gorilla/mux"
RUN go get "go.mongodb.org/mongo-driver/bson"
RUN go get "go.mongodb.org/mongo-driver/bson/primitive"
RUN go get "go.mongodb.org/mongo-driver/mongo"
RUN go get "go.mongodb.org/mongo-driver/mongo/options"

WORKDIR /go/src/socially/server
RUN CGO_ENABLED=0 BOOS=linux go build -a -installsuffix cgo -o main .

# Optimize image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=build /go/src/socially/server/main .

EXPOSE 8080

CMD ["./main"]
