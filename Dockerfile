FROM golang:1.9
ENV GOPATH=/usr/go

COPY . /usr/go/src/github.com/hitesh-goel/c1x
WORKDIR /usr/go/src/github.com/hitesh-goel/c1x
RUN go get -v
RUN go install -v

CMD ["/usr/go/bin/c1x"]
EXPOSE 8080