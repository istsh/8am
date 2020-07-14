package rpc

import (
	"google.golang.org/grpc"

	"github.com/istsh/8am/app/interface/rpc/v1.0"
	"github.com/istsh/8am/app/registry"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	v1.Apply(server, ctn)
}
