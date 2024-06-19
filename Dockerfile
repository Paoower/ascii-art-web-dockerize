FROM golang:1.22.3
WORKDIR /app
COPY . .

RUN go build -o ./server
LABEL Name=asciiartweb Version=0.0.1
CMD ["./server"]
