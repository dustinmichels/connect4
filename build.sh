# windows
GOOS=windows GOARCH=amd64 go build -o bin/connect4-amd64.exe main.go

# linux
GOOS=linux GOARCH=amd64 go build -o bin/connect4-amd64-linux main.go

# mac (amd64)
GOOS=darwin GOARCH=amd64 go build -o bin/connect4-amd64-mac main.go

# mac (arm64)
GOOS=darwin GOARCH=arm64 go build -o bin/connect4-arm64-mac main.go

# ----------------------------

# for me
go build -o bin/connect4 main.go
