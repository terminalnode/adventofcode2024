FROM golang:1.23-alpine
WORKDIR /app
COPY .. .
RUN go build -o app day01/main.go
EXPOSE 3000
CMD ["./app"]