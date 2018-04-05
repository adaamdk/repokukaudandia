package endpoint

import (
	"context"
	"time"

	pb "prjResto/gudang/grpc"
	svc "prjResto/gudang/server"
	util "prjResto/util/grpc"
	disc "prjResto/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.GudangService"
)

func NewGRPCGudangClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.GudangService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addGudangEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddGudangEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addGudangEp = retry
	}
	var updateGudangEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateGudang, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateGudangEp = retry
	}

	return GudangEndpoint{AddGudangEndpoint: addGudangEp,
		UpdateGudangEndpoint: updateGudangEp}, nil
}

func encodeAddGudangRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Gudang)
	return &pb.AddGudangReq{
		ID:           req.ID,
		NamaGudang:   req.Name,
		AlamatGudang: req.Alamat,
		LuasGudang:   req.Luas,
		Status:       req.Status,
		CreatedBy:    req.CreatedBy,
		CreatedOn:    req.CreatedOn,
		UpdatedBy:    req.UpdatedBy,
		UpdatedOn:    req.UpdatedOn,
	}, nil
}

func encodeUpdateGudangRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Gudang)
	return &pb.UpdateGudangReq{
		//proto: req.proto
		ID:           req.ID,
		NamaGudang:   req.Name,
		AlamatGudang: req.Alamat,
		LuasGudang:   req.Luas,
		Status:       req.Status,
		CreatedBy:    req.CreatedBy,
		CreatedOn:    req.CreatedOn,
		UpdatedBy:    req.UpdatedBy,
		UpdatedOn:    req.UpdatedOn,
	}, nil
}

func decodeGudangResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func makeClientAddGudangEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddGudang",
		encodeAddGudangRequest,
		decodeGudangResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddGudang")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddGudang",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateGudang(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateGudang",
		encodeUpdateGudangRequest,
		decodeGudangResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateGudang")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateGudang",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
