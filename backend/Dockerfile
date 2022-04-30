FROM golang:1.18.1-alpine3.15

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY * ./

RUN go build -o /build/webservice

EXPOSE 8080

CMD ["/build/webservice"]