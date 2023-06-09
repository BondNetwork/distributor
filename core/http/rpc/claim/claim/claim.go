// Code generated by goctl. DO NOT EDIT!
// Source: claim.proto

package claim

import (
	"context"

	"distributor/core/http/rpc/claim/types/claim"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ProofRequest  = claim.ProofRequest
	ProofResponse = claim.ProofResponse

	Claim interface {
		GetProof(ctx context.Context, in *ProofRequest, opts ...grpc.CallOption) (*ProofResponse, error)
	}

	defaultClaim struct {
		cli zrpc.Client
	}
)

func NewClaim(cli zrpc.Client) Claim {
	return &defaultClaim{
		cli: cli,
	}
}

func (m *defaultClaim) GetProof(ctx context.Context, in *ProofRequest, opts ...grpc.CallOption) (*ProofResponse, error) {
	client := claim.NewClaimClient(m.cli.Conn())
	return client.GetProof(ctx, in, opts...)
}
