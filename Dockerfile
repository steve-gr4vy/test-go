FROM golang:1.18-buster

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download

RUN make build
EXPOSE 8080
CMD ["/app/build/main"]