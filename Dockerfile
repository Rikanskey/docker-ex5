FROM golang:latest AS builder

EXPOSE 8081

WORKDIR /bapp

COPY go.* ./

RUN go mod download && go mod verify

COPY ./. ./

RUN go build -v -o . ./cmd/dockerapp

CMD ["./dockerapp"]
