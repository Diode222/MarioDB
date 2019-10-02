FROM golang:1.13
MAINTAINER Diode "diodebupt@163.com"
WORKDIR $GOPATH/src/github.com/Diode222/MarioDB
ADD . $GOPATH/src/github.com/Diode222/MarioDB
ENV GO111MODULE on
RUN go mod download && go build main.go
EXPOSE $ENV_PORT
ENTRYPOINT  ["./main"]
