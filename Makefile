gen:
	protoc --go_out=. --go_opt=Mprotos/crud_service.proto=./protos/crud_pb \
	--go-grpc_out=. --go-grpc_opt=Mprotos/crud_service.proto=./protos/crud_pb \
	protos/crud_service.proto
	protoc --go_out=. --go_opt=Mprotos/parser_service.proto=./protos/parser_pb \
	--go-grpc_out=. --go-grpc_opt=Mprotos/parser_service.proto=./protos/parser_pb \
	protos/parser_service.proto
