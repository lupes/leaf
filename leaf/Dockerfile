FROM golang:1.13

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/src/leaf

ADD . .

RUN go install -v leaf

CMD ["leaf", "server"]