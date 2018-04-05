package endpoint

import (
	"context"

	svc "prjResto/gudang/server"

	kit "github.com/go-kit/kit/endpoint"
)

// 5.
type GudangEndpoint struct {
	AddGudangEndpoint    kit.Endpoint
	UpdateGudangEndpoint kit.Endpoint
}

func NewGudangEndpoint(service svc.GudangService) GudangEndpoint {
	addGudangEp := makeAddGudangEndpoint(service)
	updateGudangEp := makeUpdateGudangEndpoint(service)

	return GudangEndpoint{AddGudangEndpoint: addGudangEp,
		UpdateGudangEndpoint: updateGudangEp,
	}
}

func makeAddGudangEndpoint(service svc.GudangService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Gudang)
		err := service.AddGudangService(ctx, req)
		return nil, err
	}
}

func makeUpdateGudangEndpoint(service svc.GudangService) kit.Endpoint {
	return func(ct context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Gudang)
		err := service.UpdateGudangService(ct, req)
		return nil, err
	}
}
