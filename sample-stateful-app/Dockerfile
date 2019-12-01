FROM golang:1.13

COPY main.go go.mod /ticket-emitter/
WORKDIR /ticket-emitter/
RUN go build -o bin .

EXPOSE 8080

ENTRYPOINT ["/ticket-emitter/bin"]

