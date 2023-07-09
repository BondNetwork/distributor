package handler

import (
	"net/http"

	"distributor/core/http/restapi/claim/internal/logic"
	"distributor/core/http/restapi/claim/internal/svc"
	"distributor/core/http/restapi/claim/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClaimHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProofRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewClaimLogic(r.Context(), svcCtx)
		resp, err := l.Claim(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
