package endpoint

import (
	"context"
	"time"

	pb "prjResto/user/grpc"
	svc "prjResto/user/server"
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
	grpcName = "grpc.UserService"
)

func NewGRPCUserClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.UserService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addUserEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddUserEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addUserEp = retry
	}
	var updateUserEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateUser, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateUserEp = retry
	}

	return UserEndpoint{AddUserEndpoint: addUserEp,
		UpdateUserEndpoint: updateUserEp}, nil
}

func encodeAddUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.User)
	return &pb.AddUserReq{
		Id:        req.ID,
		Username:  req.Username,
		Pwd:       req.Pwd,
		Idkyw:     req.IDkaryawan,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.UpdatedBy,
	}, nil
}

func encodeUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.User)
	return &pb.UpdateUserReq{
		Id:        req.ID,
		Username:  req.Username,
		Pwd:       req.Pwd,
		Idkyw:     req.IDkaryawan,
		Status:    req.Status,
		CreatedBy: req.CreatedBy,
		CreatedOn: req.CreatedOn,
		UpdatedBy: req.UpdatedBy,
		UpdatedOn: req.UpdatedOn,
	}, nil
}

func decodeUserResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func makeClientAddUserEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddUser",
		encodeAddUserRequest,
		decodeUserResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddUser")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddUser",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateUser(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateUser",
		encodeUpdateUserRequest,
		decodeUserResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateUser")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateUser",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
