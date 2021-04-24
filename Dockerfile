FROM golang:latest

RUN go build cmd/main.go

EXPOSE 8000
ENTRYPOINT ["./main"]