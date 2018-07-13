FROM golang:1.9.5

WORKDIR /go/src/app
COPY . .

RUN go build ./main.go

CMD ["./main"]
