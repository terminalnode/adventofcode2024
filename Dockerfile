FROM golang:1.23-alpine
WORKDIR /app
COPY go.mod .
COPY common common

ARG DAY
COPY solutions/day${DAY} solutions/day${DAY}
RUN go build -o app solutions/day${DAY}/main.go
EXPOSE 3000
CMD ["./app"]
