build: \
  go build -o build/server cmd/server/main.go
agent: \
  gcc -o build/agent cmd/agent/main.c
