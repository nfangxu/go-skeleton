// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 2d8d8029f5
// Version Date: 2021-02-23T10:38:53Z

package gen

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/DoNewsCode/core/contract"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/nfangxu/core-skeleton/app/proto"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC AppServer.
func MakeGRPCServer(endpoints Endpoints, options ...grpctransport.ServerOption) pb.AppServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	serverOptions = append(serverOptions, options...)
	return &grpcServer{
		// app

		getone: grpctransport.NewServer(
			endpoints.GetOneEndpoint,
			DecodeGRPCGetOneRequest,
			EncodeGRPCGetOneResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the AppServer interface
type grpcServer struct {
	getone grpctransport.Handler
}

// Methods for grpcServer to implement AppServer interface

func (s *grpcServer) GetOne(ctx context.Context, req *pb.GetOneUserRequest) (*pb.UserInfoReply, error) {
	_, rep, err := s.getone.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UserInfoReply), nil
}

// Server Decode

// DecodeGRPCGetOneRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC getone request to a user-domain getone request. Primarily useful in a server.
func DecodeGRPCGetOneRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.GetOneUserRequest)
	return req, nil
}

// Server Encode

// EncodeGRPCGetOneResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain getone response to a gRPC getone reply. Primarily useful in a server.
func EncodeGRPCGetOneResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.UserInfoReply)
	return resp, nil
}

// Helpers

func metadataToContext(ctx context.Context, md metadata.MD) context.Context {
	for k, v := range md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	ctx = context.WithValue(ctx, contract.RequestUrlKey, md.Get(":path")[0])
	ctx = context.WithValue(ctx, contract.TransportKey, "GPRCPROTOBUF")
	return ctx
}