#!/usr/bin/bash

echo "building the server..."
GOOS=linux go build -o ./bin/server ./server/server.go

echo "building the client..."
GOOS=linux go build -o ./bin/client ./cmd/client.go