module gomicro-quickstart

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/golang/protobuf v1.4.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.5.1
	google.golang.org/grpc/examples v0.0.0-20200819190100-f640ae6a4f43 // indirect
	google.golang.org/protobuf v1.24.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
