package v1

import (
	"google.golang.org/grpc"

	"github.com/istsh/8am/app/interface/rpc/v1.0/protocol"
	"github.com/istsh/8am/app/registry"
	"github.com/istsh/8am/app/usecase"
)

func Apply(server *grpc.Server, ctn *registry.Container) {
	protocol.RegisterUserServiceServer(server, NewUserService(ctn.Resolve("user-usecase").(usecase.UserUsecase)))
}
