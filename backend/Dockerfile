FROM golang:1.23.2
LABEL maintainer="Tomkoid"

WORKDIR /go/src

COPY go.mod /go/src/go.mod
COPY go.sum /go/src/go.sum

RUN go mod download
RUN go get -d -v ./...

COPY main.go /go/src/main.go
COPY ./internal /go/src/internal
COPY ./db /go/src/db
COPY ./misc /go/src/misc

RUN go build -o stravule

EXPOSE 8080
ENTRYPOINT [ "./stravule" ]
# ENTRYPOINT [ "bash" ]
