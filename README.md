## Steps to run

1. Generate/Regenerate the protobufs
  ```sh
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld/helloworld.proto
  ```

2. Run the server (The port can be ignored if you want to use the default)
  ```sh
  go run greeter_server/main.go -port 10000
  ```

3. Run the client (The addr can be ignored if the server is run with default port)
  ```sh
  go run greeter_client/main.go -addr "localhost:10000" -name "John"
  ```