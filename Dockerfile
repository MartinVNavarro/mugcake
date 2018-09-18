FROM golang:latest

WORKDIR '/go/src/mugcake'

COPY . .

RUN go get github.com/spf13/cobra
RUN go install