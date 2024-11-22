FROM golang:1.23.2
LABEL maintainer="Tomkoid"

WORKDIR /go/src

ENV CGO_ENABLED=1
RUN go env -w GOCACHE=/go-cache
RUN go env -w GOMODCACHE=/gomod-cache

COPY go.mod /go/src/go.mod
COPY go.sum /go/src/go.sum

RUN --mount=type=cache,target=/gomod-cache go mod download

COPY main.go /go/src/main.go
COPY ./internal /go/src/internal
COPY ./db /go/src/db
COPY ./misc /go/src/misc

RUN --mount=type=cache,target=/gomod-cache --mount=type=cache,target=/go-cache go build -o stravule

EXPOSE 8080
ENTRYPOINT [ "./stravule" ]
# ENTRYPOINT [ "bash" ]
