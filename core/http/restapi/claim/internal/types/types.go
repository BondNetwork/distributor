// Code generated by goctl. DO NOT EDIT.
package types

type ProofRequest struct {
	Account string `json:"account"`
}

type ProofResponse struct {
	BatchId int32    `json:"batchId"`
	Index   int32    `json:"index"`
	Account string   `json:"account"`
	Proofs  []string `json:"proofs"`
	Amount  string   `json:"amount"`
}
