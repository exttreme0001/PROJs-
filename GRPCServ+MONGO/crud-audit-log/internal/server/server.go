package server

import (
	"fmt"
	"net"

	audit "github.com/GOLANG-NINJA/crud-audit-log/pkg/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcSrv     *grpc.Server
	auditServer audit.AuditServiceServer
}

func New(auditServer audit.AuditServiceServer) *Server {
	return &Server{
		grpcSrv:     grpc.NewServer(),
		auditServer: auditServer,
	}
}

func (s *Server) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	audit.RegisterAuditServiceServer(s.grpcSrv, s.auditServer)
	reflection.Register(s.grpcSrv)

	if err := s.grpcSrv.Serve(lis); err != nil {
		return err
	}

	return nil
}
