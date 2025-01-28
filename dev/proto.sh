# remove existing go files
rm -rf proto/*.go

# generate go files
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/connect4.proto
