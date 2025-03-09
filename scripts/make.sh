#!/bin/bash
go build -o build/server ./pkg/server
gcc -Wall -Wextra -Wpedantic -static -o ./build/agent ./pkg/agent/agent.c 

