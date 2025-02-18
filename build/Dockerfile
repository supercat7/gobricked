FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
CMD ["go", "run", "cmd/server/main.go"]
