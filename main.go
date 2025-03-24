package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"

	pb "github.com/cilium/proxy/go/envoy/service/auth/v3"
)

// authServer는 pb.AuthorizationServer 인터페이스를 구현합니다.
type authServer struct {
	pb.UnimplementedAuthorizationServer
}

// Check 메서드는 Envoy/Cilium에서 전달한 CheckRequest를 받아서
// 요청의 Attributes를 로그로 출력한 후 모든 요청을 승인하는 OK 응답을 반환합니다.
func (s *authServer) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	attrList := req.GetAttributes()
	log.Printf("\nSource: %+v \nDestination: %+v \nRequests: %+v \n", attrList.GetSource(), attrList.GetDestination(), attrList.GetRequest())
	// google.rpc.Status를 이용해 OK 상태를 설정합니다.
	okStatus := &rpcstatus.Status{
		Code:    int32(codes.OK),
		Message: "OK",
	}
	
	// OkHttpResponse를 포함하는 CheckResponse를 생성합니다.
	resp := &pb.CheckResponse{
		Status: okStatus,
		HttpResponse: &pb.CheckResponse_OkResponse{
			OkResponse: &pb.OkHttpResponse{},
		},
	}
	return resp, nil
}

func main() {
	// 포트 50051에서 TCP 리스너를 생성합니다.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC 서버를 생성합니다.
	grpcServer := grpc.NewServer()
	pb.RegisterAuthorizationServer(grpcServer, &authServer{})

	log.Println("gRPC 서버가 포트 50051에서 시작되었습니다.")
	// 서버 실행
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

