package distributor

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestBalanceTree(t *testing.T) {
	balances := []Balance{
		{Account: common.HexToAddress("0x17ec8597ff92C3F44523bDc65BF0f1bE632917ff"), Amount: big.NewInt(200)},
		{Account: common.HexToAddress("0x63FC2aD3d021a4D7e64323529a55a9442C444dA0"), Amount: big.NewInt(300)},
		{Account: common.HexToAddress("0xD1D84F0e28D6fedF03c73151f98dF95139700aa7"), Amount: big.NewInt(250)},
		{Account: common.HexToAddress("0xD1D84F0e28D6fedF03c73151f98dF95139700aa8"), Amount: big.NewInt(10)},
		{Account: common.HexToAddress("0xD1D84F0e28D6fedF03c73151f98dF95139700aa1"), Amount: big.NewInt(20)},
		{Account: common.HexToAddress("0xD1D84F0e28D6fedF03c73151f98dF95139700aa2"), Amount: big.NewInt(30)},
		{Account: common.HexToAddress("0xD1D84F0e28D6fedF03c73151f98dF95139700aa3"), Amount: big.NewInt(40)},
		{Account: common.HexToAddress("0xD1D84F0e28D6fedF03c73151f98dF95139700aa0"), Amount: big.NewInt(5)},
	}
	tree, err := NewBalanceTree(balances)
	assert.Nil(t, err)

	root := tree.GetRoot()
	assert.Equal(t, root.Hex(), "0x91978bc497b90e81384a2e66872de65119d7411a95b8524a129c85eb75000bb0")

	//assert.True(t, VerifyProof(2, common.HexToAddress("0xD1D84F0e28D6fedF03c73151f98dF95139700aa7"), big.NewInt(200), p, root))

	proofs := [][]string{
		{"f7dc7c4c12b6340275991cece1c68349d534c2bcc3a9f1f3cdd81b438f9d20f3", "2a411ed78501edb696adca9e41e78d8256b61cfac45612fa0434d7cf87d916c6", "4214e9a8dc1c86b583a4997a342da1841fff86229561ba609bf38e64e3dc63c3"},
		{"bfeb956a3b705056020a3b64c540bff700c0f6c96c55c0a5fcab57124cb36f7b", "b5d2f0f0751d0bf802bd74fd4286bb964216ff4fb669fa941b0f14a81d259e89", "4214e9a8dc1c86b583a4997a342da1841fff86229561ba609bf38e64e3dc63c3"},
		{"ceaacce7533111e902cc548e961d77b23a4d8cd073c6b68ccf55c62bd47fc36b", "b5d2f0f0751d0bf802bd74fd4286bb964216ff4fb669fa941b0f14a81d259e89", "4214e9a8dc1c86b583a4997a342da1841fff86229561ba609bf38e64e3dc63c3"},
		{"b4accc54d12756efbaee1570b0da7eb3f6ae8f7602ea8d709515426dbc25e139", "0bfc4c1a894a83f15f51fe0b6e634ad5626e9b4964a75e037c64ddb39d3fe8d5", "49de6944398587bbd8a6bd3948e8f277bc8884d44f5f0d53593cf3a3a66e0e41"},
		{"2eca85cd74de3fed9b7239ade98dfea7955d2cc1419a4a5a5fa67aaeda32280d", "caf43d4c25e9f343082cc0411321b9460a7f464b75e884540a851153a5ac9814", "49de6944398587bbd8a6bd3948e8f277bc8884d44f5f0d53593cf3a3a66e0e41"},
		{"6e2f506cc671dfc955ea60fcebe70434d7de5a7f2270e2758da6ca4619475961", "caf43d4c25e9f343082cc0411321b9460a7f464b75e884540a851153a5ac9814", "49de6944398587bbd8a6bd3948e8f277bc8884d44f5f0d53593cf3a3a66e0e41"},
		{"d31de46890d4a77baeebddbd77bf73b5c626397b73ee8c69b51efe4c9a5a72fa", "2a411ed78501edb696adca9e41e78d8256b61cfac45612fa0434d7cf87d916c6", "4214e9a8dc1c86b583a4997a342da1841fff86229561ba609bf38e64e3dc63c3"},
		{"9430e623c11ecf8bff61cc6af8ba53e842fa09336d68f7e4cc5ec8daf9caa64c", "0bfc4c1a894a83f15f51fe0b6e634ad5626e9b4964a75e037c64ddb39d3fe8d5", "49de6944398587bbd8a6bd3948e8f277bc8884d44f5f0d53593cf3a3a66e0e41"},
	}
	for idx, balance := range balances {
		p := make(Elements, 0)
		for _, s := range proofs[idx] {
			p = append(p, common.HexToHash(s))
		}
		assert.True(t, VerifyProof(idx, balance.Account, balance.Amount, p, root))
	}

	info, err := ParseBalanceMap(balances)

	data, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	//assert.Nil(t, err)
	//assert.Equal(t, info.MerkleRoot.Hex(), "0x2ec9c2fc2a55df417ba88ecd833f165fa3c5941772ebaf8c5f4debe33f4d1b12")
	//assert.Equal(t, info.TokenTotal, "0x2ee")
}
