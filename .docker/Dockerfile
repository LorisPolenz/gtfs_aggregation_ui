# Specifies a parent image
FROM golang:1.23.3-bullseye AS builder
WORKDIR /build
COPY . .
RUN apt-get update && apt-get install build-essential -y
RUN go mod download
RUN go mod tidy 
RUN go mod verify 
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-s -w" -o main .

# copy to lean production image
FROM debian:bullseye-slim
COPY --from=builder /build/main /app/
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["/app/main"]


