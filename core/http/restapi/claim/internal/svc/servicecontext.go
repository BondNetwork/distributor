package svc

import (
	"distributor/core/http/restapi/claim/internal/config"
	"distributor/core/http/rpc/claim/claim"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	ClaimRpc claim.Claim
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		ClaimRpc: claim.NewClaim(zrpc.MustNewClient(c.ClaimRpc)),
	}
}
