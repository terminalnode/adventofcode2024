FROM golang:1.23-alpine
WORKDIR /app
COPY .. .
RUN go build -o app day03/main.go
EXPOSE 3000
CMD ["./app"]