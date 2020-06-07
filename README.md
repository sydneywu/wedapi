### generate pb file
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative api/photographer/photographer.proto