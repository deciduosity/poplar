package rpc

import (
	"github.com/deciduosity/poplar"
	"github.com/deciduosity/poplar/rpc/internal"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func AttachService(registry *poplar.RecorderRegistry, s *grpc.Server) error {
	return errors.WithStack(internal.AttachService(registry, s))
}
