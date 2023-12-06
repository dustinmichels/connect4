# windows
GOOS=windows GOARCH=amd64 go build -o bin/connect4-amd64.exe main.go

# linux
GOOS=linux GOARCH=amd64 go build -o bin/connect4-amd64-linux main.go
