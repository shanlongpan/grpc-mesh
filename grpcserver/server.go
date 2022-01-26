/**
* @Author:Tristan
* @Date: 2022/1/18 3:11 下午
 */

package grpcserver

import (
	"github.com/shanlongpan/grpc-mesh/grpc/user"
	"github.com/shanlongpan/grpc-mesh/impl"
	"google.golang.org/grpc/xds"
	"log"
	"net"
	"sync"
)

func Serve(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("[User] GRPC failed to listen: %v", err)
	}

	//s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	//	grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(opentracing.GlobalTracer())),
	//	grpc_recovery.UnaryServerInterceptor(),
	//)))

	s := xds.NewGRPCServer()
	user.RegisterUserServiceServer(s, &impl.Server{})

	log.Printf("[User] Serving GRPC on localhost:%s ...", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[User] GRPC failed to serve: %v", err)
	}
}
