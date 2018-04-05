package endpoint

import (
	"context"

	pb "prjResto/user/grpc"
	scv "prjResto/user/server"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcUserServer struct {
	addUser    grpctransport.Handler
	updateUser grpctransport.Handler
}

func NewGRPCUserServer(endpoints UserEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.UserServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcUserServer{
		addUser: grpctransport.NewServer(endpoints.AddUserEndpoint,
			decodeAddUserRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddUser", logger)))...),

		updateUser: grpctransport.NewServer(endpoints.UpdateUserEndpoint,
			decodeUpdateUserRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateUser", logger)))...),
	}
}

func decodeAddUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddUserReq)
	return scv.User{
		ID:         req.GetId(),
		Username:   req.GetUsername(),
		Pwd:        req.GetPwd(),
		IDkaryawan: req.GetIdkyw(),
		Status:     req.GetStatus(),
		CreatedBy:  req.GetCreatedBy(),
		CreatedOn:  req.GetCreatedOn(),
		UpdatedBy:  req.GetUpdatedBy(),
		UpdatedOn:  req.GetUpdatedOn()}, nil
}

func decodeUpdateUserRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateUserReq)
	return scv.User{
		ID:         req.Id,
		Username:   req.Username,
		Pwd:        req.Pwd,
		IDkaryawan: req.Idkyw,
		Status:     req.Status,
		CreatedBy:  req.CreatedBy,
		CreatedOn:  req.CreatedOn,
		UpdatedBy:  req.UpdatedBy,
		UpdatedOn:  req.UpdatedOn}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func (s *grpcUserServer) AddUser(ctx oldcontext.Context, shift *pb.AddUserReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addUser.ServeGRPC(ctx, shift)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcUserServer) UpdateUser(ctx oldcontext.Context, js *pb.UpdateUserReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateUser.ServeGRPC(ctx, js)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
