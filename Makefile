
default:
	dep ensure
	protoc -I. --go_out=plugins=grpc:. model/*.proto
	go run cmd/main.go
