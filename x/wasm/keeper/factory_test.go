package keeper

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

type FactoryInitMsg struct {
	// Contract can instantiate itself through a boolean flag
	Factory bool `json:"factory"`
	// CodeID is the child contract to instantiate
	CodeID uint64 `json:"code_id"`

	Controller bool `json:"controller"`
}

func (m FactoryInitMsg) GetBytes(t testing.TB) []byte {
	initMsgBz, err := json.Marshal(m)
	require.NoError(t, err)
	return initMsgBz
}

func TestFactoryContracts(t *testing.T) {
	ctx, keepers := CreateTestInput(t, false, AvailableCapabilities)
	contract := StoreFactoryContract(t, ctx, keepers)
	// contract2 := StoreFactoryContract(t, ctx, keepers)

	initMsgBz := FactoryInitMsg{
		Factory:    true,
		CodeID:     contract.CodeID,
		Controller: false,
	}.GetBytes(t)
	initialAmount := sdk.NewCoins(sdk.NewInt64Coin("denom", 100))

	adminAddr := contract.CreatorAddr
	label := "demo contract to query"
	_, _, err := keepers.ContractKeeper.Instantiate(ctx, contract.CodeID,
		contract.CreatorAddr, adminAddr, initMsgBz, label, initialAmount)
	require.NoError(t, err)

}
