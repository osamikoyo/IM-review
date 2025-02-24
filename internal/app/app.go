package app

import (
	"fmt"
	"net"

	"github.com/osamikoyo/IM-review/internal/config"
	"github.com/osamikoyo/IM-review/internal/data"
	"github.com/osamikoyo/IM-review/internal/server"
	"github.com/osamikoyo/IM-review/pkg/loger"
	"github.com/osamikoyo/IM-review/pkg/proto/pb"
	"google.golang.org/grpc"
)

type App struct{
	gRPC *grpc.Server
	cfg *config.Config
	loger loger.Logger
	server *server.Server
}

func Init() (*App, error) {
	cfg, err := config.Load("config.yml")
	if err != nil{
		return nil, err
	}

	st, err := data.New(cfg)
	if err != nil{
		return nil,err
	}

	grpcServer := grpc.NewServer()
	server := &server.Server{
		Storage: st,
	}
	return &App{
		gRPC: grpcServer,
		server: server,
		cfg: cfg,
		loger: loger.New(),
	}, nil
}

func (a *App) Run() error {
	a.loger.Info().Msg("starting review service...")

	pb.RegisterReviewServiceServer(a.gRPC, a.server)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port))
	if err != nil{
		return err
	}

	a.loger.Info().Str("addr", lis.Addr().String()).Msg("server started!")
	return a.gRPC.Serve(lis)
}