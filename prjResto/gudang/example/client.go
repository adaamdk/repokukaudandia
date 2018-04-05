package main

import (
	"context"
	"time"

	cli "prjResto/gudang/endpoint"
	svc "prjResto/gudang/server"
	opt "prjResto/util/grpc"
	util "prjResto/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCGudangClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add

	client.AddGudangService(context.Background(), svc.Gudang{ID: "003", Name: "Gudang-03", Alamat: "Jl. Raflesia", Luas: "75", CreatedBy: "Adam", UpdatedBy: "Adam"})

	//Update
	client.UpdateGudangService(context.Background(), svc.Gudang{ID: "004", Name: "Gudang-04", Alamat: "Jl. Mawar", Luas: "50", CreatedBy: "Adam", UpdatedBy: "Adam"})

}
