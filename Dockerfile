FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download 
RUN go build -o bin/bh cmd/api/main.go

EXPOSE 8080

CMD ["./bin/bh"]