FROM golang:latest

ADD . .

RUN go build

EXPOSE 8081

CMD ["./go-web-hello-world"]
