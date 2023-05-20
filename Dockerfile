FROM golang:1.20

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go

RUN chmod a+x main

CMD ["./main"]