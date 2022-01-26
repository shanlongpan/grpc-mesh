/**
* @Author:Tristan
* @Date: 2022/1/18 3:13 下午
 */

package impl

import (
	"context"
	"fmt"
	"github.com/shanlongpan/grpc-mesh/grpc/user"
	"github.com/shanlongpan/grpc-mesh/lib"
	"log"
)

type Server struct {
	user.UnimplementedUserServiceServer
}

func (s *Server) Create(ctx context.Context, in *user.CreateUserReq) (*user.CreateUserResp, error) {

	log.Printf("[User] Create Req: %v", in.GetUser())
	id, err := lib.DefaultSnow.NextID()
	if err != nil {
		log.Fatalln(err)
	}
	r := &user.CreateUserResp{Id: int64(id)}

	log.Printf("[User] Create Res: %v", r.GetId())

	return r, nil
}

func (s *Server) Read(ctx context.Context, in *user.ReadUserReq) (*user.ReadUserResp, error) {

	log.Printf("[User] Read Req: %v", in.GetId())

	r := &user.ReadUserResp{User: &user.User{Id: in.GetId(), Name: fmt.Sprintf("user%d", lib.RandomInt(1000000)), Email: fmt.Sprintf("%s@163.com", lib.GetRandomString2(10))}}

	log.Printf("[User] Read Res: %v", r.GetUser())

	return r, nil
}
