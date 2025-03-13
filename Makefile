PROTO_DIR=api/proto
OUT_DIR=./pb

build-proto:
	protoc --go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_DIR)/*.proto

build:
	go build -o bin/server ./cmd/server/main.go

run:
	go run ./cmd/server/main.go

clean:
	rm -rf ./bin

proto: build-proto