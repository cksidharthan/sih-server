# Build stage
FROM golang:1.21.5-alpine3.19 AS build
WORKDIR /app
COPY . .
RUN go build -o main .

# Run stage
FROM scratch
WORKDIR /app
EXPOSE 9090
COPY --from=build /app/main .
CMD ["/app/main"]