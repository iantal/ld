.PHONY: protos

protos:
	 protoc -I protos/ protos/ld.proto --go_out=plugins=grpc:protos/ld