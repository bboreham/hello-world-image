FROM golang:1.9-alpine3.6
WORKDIR /home/weave
COPY server.go logo.png index.template ./
RUN go build -o server .
ENTRYPOINT ["./server"]
EXPOSE 80/tcp
