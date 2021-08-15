module service_1

go 1.16

replace common => ../common

require (
	common v0.0.0-00010101000000-000000000000
	github.com/emicklei/go-restful/v3 v3.5.2
	github.com/json-iterator/go v1.1.11 // indirect
	google.golang.org/grpc v1.39.1
	google.golang.org/grpc/examples v0.0.0-20210512000516-62adda2ece5e
)
