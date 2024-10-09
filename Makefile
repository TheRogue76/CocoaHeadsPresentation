generate_go_defs:
	protoc --proto_path=ServiceDefinitions ServiceDefinitions/*.proto --go_out=. --go-grpc_out=.
generate_swift_defs:
	protoc --proto_path=ServiceDefinitions ServiceDefinitions/*.proto --swift_opt=Visibility=Public --swift_out=SwiftGrpcGenerated/Sources --grpc-swift_opt=Client=true,Server=false,Visibility=Public --grpc-swift_out=SwiftGrpcGenerated/Sources