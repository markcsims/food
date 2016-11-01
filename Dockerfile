FROM golang:1.7

RUN go get -u github.com/alecthomas/gometalinter
RUN go get -u github.com/kardianos/govendor
RUN gometalinter --install

WORKDIR /go/src/github.com/markcsims/bbcfood

ADD . /go/src/github.com/markcsims/bbcfood

CMD ./build-app.sh && bbcfood

EXPOSE 8080
