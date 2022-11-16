FROM golang as builder
ENV GO111MODULE=on
WORKDIR /app
COPY src /app/

# RUN CGO=0 GOOS=linux GOARCH=amd64 go build -o bin/gin_api
RUN go build -o bin/gin_api

# Stage 2: build the image
FROM alpine
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /app/
COPY --from=builder  /app/bin/gin_api .

EXPOSE 8080
# Run the Go Gin binary.
CMD ["./gin_api"]