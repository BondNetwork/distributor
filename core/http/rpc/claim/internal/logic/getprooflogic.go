package logic

import (
	"context"

	"distributor/core/http/rpc/claim/internal/svc"
	"distributor/core/http/rpc/claim/types/claim"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProofLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProofLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProofLogic {
	return &GetProofLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProofLogic) GetProof(in *claim.ProofRequest) (*claim.ProofResponse, error) {
	// todo: add your logic here and delete this line
	return &claim.ProofResponse{
		BatchId: 123,
		Index:   1,
		Account: "0xed12121212",
		Proofs:  []string{"0x123", "0x124"},
		Amount:  "23232323232323232323",
	}, nil
}
