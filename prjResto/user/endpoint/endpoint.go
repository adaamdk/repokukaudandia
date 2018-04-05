package endpoint

import (
	"context"

	svc "prjResto/user/server"

	kit "github.com/go-kit/kit/endpoint"
)

// 5.
type UserEndpoint struct {
	AddUserEndpoint    kit.Endpoint
	UpdateUserEndpoint kit.Endpoint
}

func NewUserEndpoint(service svc.UserService) UserEndpoint {
	addUserEp := makeAddUserEndpoint(service)
	updateUserEp := makeUpdateUserEndpoint(service)

	return UserEndpoint{AddUserEndpoint: addUserEp,
		UpdateUserEndpoint: updateUserEp,
	}
}

func makeAddUserEndpoint(service svc.UserService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.User)
		err := service.AddUserService(ctx, req)
		return nil, err
	}
}

func makeUpdateUserEndpoint(service svc.UserService) kit.Endpoint {
	return func(ct context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.User)
		err := service.UpdateUserService(ct, req)
		return nil, err
	}
}
