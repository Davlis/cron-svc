FROM golang:latest AS build-env

RUN mkdir -p $GOPATH/src/github.com/davlis/services/cron
WORKDIR $GOPATH/src/github.com/davlis/services/cron
COPY . ./
COPY . /usr/local/go/src/github.com/davlis/cron/

ENV GIT_SSL_NO_VERIFY=1

RUN go get -u github.com/golang/dep/...

RUN go build -o main .

EXPOSE 3000
CMD ["./main"]
