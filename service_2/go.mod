module service_2

go 1.16

replace common => ../common

require (
	common v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.0
	google.golang.org/grpc v1.39.1
	google.golang.org/grpc/examples v0.0.0-20210512000516-62adda2ece5e
)
