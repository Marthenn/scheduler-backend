FROM golang:1.20

WORKDIR /app

COPY . .
RUN go build -o app ./cmd/app

EXPOSE 8081

CMD ["./app"]