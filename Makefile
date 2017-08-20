test:
	go1.9rc2 test -race ./...

proto:
	protoc -I plugin/proto/ plugin/proto/plugin.proto --go_out=plugins=grpc:plugin/proto
	protoc -I server/proto/ server/proto/server.proto --go_out=plugins=grpc:server/proto
