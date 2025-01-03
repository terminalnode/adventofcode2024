FROM golang:1.23-alpine AS golang-with-curl
RUN apk --no-cache add curl

FROM golang-with-curl AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM deps AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY common common

ARG DAY
COPY solutions/day${DAY} solutions/day${DAY}
RUN go build -o app ./solutions/day${DAY}

FROM golang-with-curl
WORKDIR /app
COPY --from=build /app/app .

EXPOSE 3000
EXPOSE 50051
HEALTHCHECK \
    --start-interval=10s \
    --interval=10s \
    --timeout=3s \
    CMD curl -f http://127.0.0.1:3000/health

CMD ["./app"]
