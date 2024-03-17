FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build -o bh main.go
CMD ["./main"]