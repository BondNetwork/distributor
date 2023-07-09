package logic

import (
	"context"
	"distributor/core/http/rpc/claim/claim"

	"distributor/core/http/restapi/claim/internal/svc"
	"distributor/core/http/restapi/claim/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClaimLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClaimLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClaimLogic {
	return &ClaimLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClaimLogic) Claim(req *types.ProofRequest) (resp *types.ProofResponse, err error) {
	// todo: add your logic here and delete this line
	res, error := l.svcCtx.ClaimRpc.GetProof(l.ctx, &claim.ProofRequest{Account: "0xabc"})
	if error != nil {
		return nil, error
	}
	return &types.ProofResponse{
		BatchId: res.BatchId,
		Index:   res.Index,
		Account: res.Account,
		Proofs:  res.Proofs,
		Amount:  res.Amount,
	}, nil
}
