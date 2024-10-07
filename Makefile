generate_go_defs:
    protoc --proto_path=api-def api-def/*.proto --go_out=server-def --go-grpc_out=server-def
generate_swift_defs:
	protoc --proto_path=api-def api-def/*.proto --swift_opt=Visibility=Public --swift_out=ios-def/Sources/ios-def --grpc-swift_opt=Client=true,Server=false,Visibility=Public --grpc-swift_out=ios-def/Sources/ios-def