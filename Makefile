proto:
	protoc --go_out=. --go-grpc_out=. pkg/auth/pb/auth.proto
	protoc --go_out=. --go-grpc_out=. pkg/product/pb/product.proto
	protoc --go_out=. --go-grpc_out=. pkg/order/pb/order.proto

server:
	 go run cmd/main.go
